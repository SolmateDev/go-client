package client

import (
	"context"
	"crypto/tls"
	"errors"
	"os"
	"time"

	pbauth "github.com/SolmateDev/go-client/proto/auth"
	pbtstr "github.com/SolmateDev/go-client/proto/solana/tester"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
)

type Configuration struct {
	Hostname string
	ApiKey   string
}

type Client struct {
	config  *Configuration
	jwtKey  string
	conn    *grpc.ClientConn
	Session pbauth.SessionClient
	Solana  *Solana
}

type Solana struct {
	Sandbox pbtstr.ValidatorEnvClient
}

func ConfigFromEnv() (*Configuration, error) {
	config := new(Configuration)
	_, present := os.LookupEnv("TEST")
	if present {
		config.Hostname = "grpc.test.solmate.dev"
	} else {
		config.Hostname = "grpc.solmate.dev"
	}

	config.ApiKey, present = os.LookupEnv("API_KEY")
	if !present {
		return nil, errors.New("$API_KEY not set")
	}
	return config, nil
}

// Connect from the outside to the inside using this client.  Set the JWT key for each grpc call.  The JWT contains the user id, from which the book id (tenant id) can be looked up from the database.
func Create(ctx context.Context, config *Configuration) (*Client, error) {
	var err error

	if config == nil {
		config, err = ConfigFromEnv()
		if err != nil {
			return nil, err
		}
	}

	e1 := new(Client)
	e1.conn, err = e1.connect(ctx)
	if err != nil {
		return nil, err
	}

	e1.config = config
	e1.Session = pbauth.NewSessionClient(e1.conn)
	s := new(Solana)
	e1.Solana = s
	s.Sandbox = pbtstr.NewValidatorEnvClient(e1.conn)

	err = e1.login(ctx)
	if err != nil {
		return nil, err
	}

	return e1, nil
}

func (e1 *Client) connect(ctx context.Context) (*grpc.ClientConn, error) {
	var kacp = keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             2 * time.Second,  // wait 1 second for ping back
		PermitWithoutStream: true,             // send pings even without active streams
	}
	return grpc.DialContext(
		ctx,
		e1.config.Hostname,
		grpc.WithTransportCredentials(credentials.NewTLS(
			&tls.Config{
				ServerName: e1.config.Hostname,
			},
		)),
		grpc.WithKeepaliveParams(kacp),
	)
}

func (e1 *Client) login(ctx context.Context) error {
	var err error

	resp, err := e1.Session.CreateByApiKey(ctx, &pbauth.ApiKeySessionRequest{Key: e1.config.ApiKey})
	if err != nil {
		return err
	}
	e1.jwtKey = resp.GetJwt()

	return nil
}

func (e1 *Client) Ctx(ctx context.Context) context.Context {
	md := metadata.New(map[string]string{"JWT": e1.jwtKey})
	return metadata.NewOutgoingContext(ctx, md)
}
