// This package creates a Solana Test validator accessible via HTTP.
package sandbox

import (
	"context"
	"errors"
	"io"

	clt "github.com/SolmateDev/go-client/client"
	pbsol "github.com/SolmateDev/go-client/proto/solana"
	pbtstr "github.com/SolmateDev/go-client/proto/solana/tester"
	log "github.com/sirupsen/logrus"
)

func Run(ctx context.Context) error {

	client, err := clt.Create(ctx, nil)
	if err != nil {
		return err
	}

	stream, err := client.Solana.Sandbox.CreatePrivateCluster(client.Ctx(ctx), &pbsol.PrivateClusterCreateRequest{
		Count: 3,
	})
	if err != nil {
		return err
	}
	defer stream.CloseSend()

	doneC := ctx.Done()
	progressC := make(chan *pbsol.Progress, 10)
	clusterC := make(chan *pbsol.PrivateCluster, 10)
	validatorC := make(chan *pbsol.PrivateValidator, 10)
	logC := make(chan *pbsol.TransactionLog, 10)
	errorC := make(chan error, 1)

	// once the stream is closed or ctx.Done() is reached, the sandbox on the Solmate side will be terminated
	go readStream(stream, errorC, progressC, clusterC, validatorC, logC)

out:
	for {
		select {
		case <-doneC:
			break out
		case err = <-errorC:
			break out
		case p := <-progressC:
			log.Debugf("progress: %+v", p)
		case c := <-clusterC:
			log.Debugf("cluster: %+v", c)
		case v := <-validatorC:
			log.Debugf("validator: %+v", v)
		case l := <-logC:
			log.Debugf("log: %+v", l)
		}
	}

	if err != nil {
		return err
	}

	return nil
}

func readStream(stream pbtstr.ValidatorEnv_CreatePrivateClusterClient, errorC chan<- error, progressC chan<- *pbsol.Progress, clusterC chan<- *pbsol.PrivateCluster, validatorC chan<- *pbsol.PrivateValidator, logC chan<- *pbsol.TransactionLog) {
	var err error
	var msg *pbsol.PrivateClusterCreateResponse
out:
	for {
		msg, err = stream.Recv()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break out
		}

		log.Debugf("msg=%+v", msg)
		switch msg.Response.(type) {
		case *pbsol.PrivateClusterCreateResponse_Empty:
			// do nothing, just for pinging
			log.Debug("received ping")
		case *pbsol.PrivateClusterCreateResponse_Progress:
			x, ok := msg.Response.(*pbsol.PrivateClusterCreateResponse_Progress)
			if !ok {
				err = errors.New("bad progress")
			}
			if x.Progress == nil {
				err = errors.New("blank progress")
			}

		case *pbsol.PrivateClusterCreateResponse_Cluster:
			x, ok := msg.Response.(*pbsol.PrivateClusterCreateResponse_Cluster)
			if !ok {
				err = errors.New("bad cluster")
				break out
			}
			if x.Cluster == nil {
				err = errors.New("blank cluster")
				break out
			}
			clusterC <- x.Cluster
		case *pbsol.PrivateClusterCreateResponse_Log:
			x, ok := msg.Response.(*pbsol.PrivateClusterCreateResponse_Log)
			if !ok {
				err = errors.New("bad log")
				break out
			}
			if x.Log == nil {
				err = errors.New("blank log")
				break out
			}
			logC <- x.Log
		case *pbsol.PrivateClusterCreateResponse_Validator:
			x, ok := msg.Response.(*pbsol.PrivateClusterCreateResponse_Validator)
			if !ok {
				err = errors.New("bad validator")
				break out
			}
			if x.Validator == nil {
				err = errors.New("blank validator")
				break out
			}
			validatorC <- x.Validator
		}
	}
	errorC <- err
}
