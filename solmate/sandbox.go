package main

import (
	"errors"
	"io"
	"os"

	pbsol "github.com/SolmateDev/go-client/proto/solana"
	pbtstr "github.com/SolmateDev/go-client/proto/solana/tester"
	log "github.com/sirupsen/logrus"
)

type Sandbox struct {
	Count  uint32 `arg name:"count" short:"n" help:"how many nodes will be in the private cluster"`
	Output string `option name:"output" short:"o" help:"where will the result of the create command go."`
}

func (r *Sandbox) Run(kongCtx *CLIContext) error {
	li := kongCtx.LocalInfo
	ctx := li.ctx
	err := li.auth()
	if err != nil {
		return err
	}
	client := li.client
	stream, err := client.Solana.Sandbox.CreatePrivateCluster(client.Ctx(ctx), &pbsol.PrivateClusterCreateRequest{
		Count: r.Count,
	})
	if err != nil {
		return err
	}
	doneC := ctx.Done()
	progressC := make(chan *pbsol.Progress, 10)
	clusterC := make(chan *pbsol.PrivateCluster, 10)
	validatorC := make(chan *pbsol.PrivateValidator, 10)
	logC := make(chan *pbsol.TransactionLog, 10)
	errorC := make(chan error, 1)

	splitter := "----------"

	var logger *os.File
	if len(r.Output) == 0 {
		log.Debug("output to stdout")
		logger = os.Stdout
	} else if r.Output == "-" {
		log.Debug("output to stdout")
		logger = os.Stdout
	} else {
		log.Debugf("output to %s", r.Output)
		logger, err = os.Open(r.Output)
		if err != nil {
			return err
		}
	}
	logger.Write([]byte("\n" + splitter + "\n"))
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
			err = printEvent(logger, splitter, p)
			if err != nil {
				break out
			}
		case c := <-clusterC:
			err = printEvent(logger, splitter, c)
			if err != nil {
				break out
			}
		case v := <-validatorC:
			err = printEvent(logger, splitter, v)
			if err != nil {
				break out
			}
		case l := <-logC:
			err = printEvent(logger, splitter, l)
			if err != nil {
				break out
			}
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
			log.Debug("received progress")
			x, ok := msg.Response.(*pbsol.PrivateClusterCreateResponse_Progress)
			if !ok {
				err = errors.New("bad progress")
			}
			if x.Progress == nil {
				err = errors.New("blank progress")
			}
			progressC <- x.Progress
		case *pbsol.PrivateClusterCreateResponse_Cluster:
			log.Debug("received cluster")
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
			log.Debug("received log")
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
			log.Debug("received validator")
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
