package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	clt "github.com/SolmateDev/go-client/client"
	pba "github.com/SolmateDev/go-client/proto/auth"
	"github.com/alecthomas/kong"
	log "github.com/sirupsen/logrus"
)

type Configuration struct {
	Host   string `json:"host"`
	Port   uint16 `json:"port"`
	ApiKey string `json:"key"`
}

func ConfigFromEnv() (*Configuration, error) {
	c := new(Configuration)
	var present bool
	var err error

	c.Host, present = os.LookupEnv("SOLMATE_HOST")
	if !present {
		return nil, errors.New("no host")
	}
	portStr, present := os.LookupEnv("SOLMATE_PORT")
	if !present {
		return nil, errors.New("no port")
	}
	portNum, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil {
		return nil, err
	}
	c.Port = uint16(portNum)

	c.ApiKey, present = os.LookupEnv("API_KEY")
	if !present {
		return nil, errors.New("no api key")
	}

	return c, nil
}

func ConfigFromFile(filePath string) (*Configuration, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	c := new(Configuration)
	err = json.NewDecoder(file).Decode(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

type CLIContext struct {
	LocalInfo *LocalInfo
	Ctx       context.Context
}

type debugFlag bool

type GrpcUrl string

var cli struct {
	Verbose debugFlag `help:"Set logging to verbose." short:"v" default:"false"`
	Grpc    GrpcUrl   `option name:"grpc" help:"URL to the Grpc Endpoint (default is grpc.solmate.dev)"`
	Sandbox Sandbox   `cmd name:"sandbox" help:"Create private Solana cluster"`
}

type LocalInfo struct {
	ctx     context.Context
	client  *clt.Client
	session *pba.ApiKeySessionResponse
}

func (d debugFlag) AfterApply(li *LocalInfo) error {
	if d {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	return nil
}

func (li *LocalInfo) auth() error {
	var err error
	li.client, err = clt.Create(li.ctx, nil)
	if err != nil {
		return err
	}
	apikey, present := os.LookupEnv("API_KEY")
	if !present {
		return errors.New("no api key")
	}
	li.session, err = li.client.Session.CreateByApiKey(li.ctx, &pba.ApiKeySessionRequest{Key: apikey})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	signalC := make(chan os.Signal, 1)
	signal.Notify(signalC, syscall.SIGTERM, syscall.SIGINT)
	ctx, cancel := context.WithCancel(context.Background())
	go loopSignal(ctx, cancel, signalC)
	li := &LocalInfo{ctx: ctx}
	kongCtx := kong.Parse(&cli, kong.Bind(li))
	err := kongCtx.Run(&CLIContext{Ctx: ctx, LocalInfo: li})
	kongCtx.FatalIfErrorf(err)
}

func loopSignal(ctx context.Context, cancel context.CancelFunc, signalC <-chan os.Signal) {
	defer cancel()
	doneC := ctx.Done()
	select {
	case <-doneC:
	case s := <-signalC:
		os.Stderr.WriteString(fmt.Sprintf("%s\n", s.String()))
	}
}

func printEvent(logger *os.File, splitter string, event interface{}) error {
	data, err := json.MarshalIndent(event, "", "    ")
	if err != nil {
		return err
	}
	logger.Write([]byte(fmt.Sprintf("\n%s[%s]\n", splitter, time.Now().Format("2006-01-02T15:04:05Z07:00"))))
	logger.Write(data)

	//err = logger.Sync()
	//if err != nil {
	//	return err
	//}
	return nil
}
