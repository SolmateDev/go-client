// This package creates a Solana Test validator accessible via HTTP.
package sandbox

import (
	"context"
	"errors"
	"net/http"

	clt "github.com/SolmateDev/go-client/client"
	pbb "github.com/SolmateDev/go-client/proto/solana/basic"
	pbtstr "github.com/SolmateDev/go-client/proto/solana/tester"
	sgo "github.com/SolmateDev/solana-go"
	sgorpc "github.com/SolmateDev/solana-go/rpc"
	log "github.com/sirupsen/logrus"
)

func Run(ctx context.Context) error {
	mint, err := sgo.NewRandomPrivateKey()
	if err != nil {
		return err
	}

	client, err := clt.Create(ctx, nil)
	if err != nil {
		return err
	}

	stream, err := client.Solana.Sandbox.Create(client.Ctx(ctx), &pbtstr.CreateRequest{
		MintAddress: &pbb.Pubkey{
			Data: mint.PublicKey().Bytes(),
		},
		TargetLamportsPerSignature: 200,
	})
	if err != nil {
		return err
	}
	defer stream.CloseSend()
	urlC := make(chan urlAns, 1)
	// once the stream is closed or ctx.Done() is reached, the sandbox on the Solmate side will be terminated
	go readStream(stream, urlC)
	u := <-urlC
	if u.err != nil {
		return u.err
	}
	// the sandbox id authorizes the request to be proxied to the Solana test validator sandbox
	customHeaders := http.Header{"Sandbox": []string{u.id}}
	rpcClient := sgorpc.NewWithHeaders(u.rpcUrl, customHeaders)

	// get the epoch as a test to see if the proxy connection works
	epoch, err := rpcClient.GetEpochInfo(ctx, sgorpc.CommitmentProcessed)
	if err != nil {
		return err
	}
	log.Debugf("epoch=%+v", epoch)

	return nil
}

type urlAns struct {
	err    error
	id     string
	rpcUrl string
	wsUrl  string
}

func readStream(stream pbtstr.ValidatorEnv_CreateClient, urlC chan<- urlAns) {
	firstPass := true
	for {
		msg, err := stream.Recv()
		if err != nil {
			if firstPass {
				urlC <- urlAns{err: err}
			}
			break
		}
		firstPass = false
		log.Debugf("msg=%+v", msg)
		switch msg.Event.(type) {
		case *pbtstr.CreateResponse_Info:
			x, ok := msg.Event.(*pbtstr.CreateResponse_Info)
			if ok {
				urlC <- urlAns{
					id:     x.Info.Id,
					rpcUrl: x.Info.RpcUrl,
					wsUrl:  x.Info.WsUrl,
					err:    nil,
				}
			} else {
				urlC <- urlAns{err: errors.New("bad type")}
			}
		}
	}
}
