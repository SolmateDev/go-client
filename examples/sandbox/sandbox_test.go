package sandbox_test

import (
	"context"
	"os"
	"testing"
	"time"

	sb "github.com/SolmateDev/go-client/examples/sandbox"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func TestSandbox(t *testing.T) {
	var err error
	_, present := os.LookupEnv("API_KEY")
	if !present {
		err = godotenv.Load("../../.env")
	}
	log.SetLevel(log.DebugLevel)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)

	t.Cleanup(func() {
		cancel()
	})

	err = sb.Run(ctx)
	if err != nil {
		t.Fatal(err)
	}

}
