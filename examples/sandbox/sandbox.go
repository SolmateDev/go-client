package sandbox

import (
	"context"
	"errors"
	"net/http"

	clt "github.com/SolmateDev/go-client/client"
	pbb "github.com/SolmateDev/go-client/proto/solana/basic"
	pbtstr "github.com/SolmateDev/go-client/proto/solana/tester"
	sgo "github.com/gagliardetto/solana-go"
	sgorpc "github.com/gagliardetto/solana-go/rpc"
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

	stream, err := client.Solana.Sandbox.Create(ctx, &pbtstr.CreateRequest{
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
	go readStream(stream, urlC)
	u := <-urlC
	if u.err != nil {
		return u.err
	}
	customHeaders := http.Header{"Sandbox": []string{u.id}}
	rpcClient := sgorpc.NewWithHeaders(u.rpcUrl, customHeaders)
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
