package trading

import (
	"context"
	"encoding/json"
	"os"

	ie "github.com/SolmateDev/go-client/examples"
	sgo "github.com/gagliardetto/solana-go"
)

type Configuration struct {
	PrivateKeyFilePath string `json:"key_file"`
	Cluster            string `json:"cluster"`
}

// Read in the configuration file from the file path indicated by the environmental file CONFIG_FILE.
func ConfigFromEnv() (*Configuration, error) {
	configData, err := os.ReadFile(os.Getenv("CONFIG_FILE"))
	if err != nil {
		return nil, err
	}
	config := new(Configuration)
	err = json.Unmarshal(configData, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

type external struct {
	internalC   chan<- func(*internal)
	private_key sgo.PrivateKey
	public_key  sgo.PublicKey
}

func New(config *Configuration, ctx context.Context) (ie.Trading, error) {
	private_key, err := sgo.PrivateKeyFromSolanaKeygenFile(config.PrivateKeyFilePath)
	if err != nil {
		return nil, err
	}
	public_key := private_key.PublicKey()
	internalC := make(chan func(*internal), 10)
	go loopInternal(ctx, internalC)
	return external{internalC: internalC, private_key: private_key, public_key: public_key}, nil
}

func (e1 external) PublicKey() sgo.PublicKey {
	return e1.public_key
}

func (e1 external) CloseSignal() <-chan error {
	signalC := make(chan error, 1)
	e1.internalC <- func(in *internal) {
		in.closeSignalCList = append(in.closeSignalCList, signalC)
	}
	return signalC
}
