package sandbox_test

import (
	"context"
	"os"
	"testing"

	sb "github.com/SolmateDev/go-client/examples/sandbox"
	"github.com/joho/godotenv"
)

func TestSandbox(t *testing.T) {
	var err error
	_, present := os.LookupEnv("API_KEY")
	if !present {
		err = godotenv.Load("../../.env")
	}
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())

	t.Cleanup(func() {
		cancel()
	})

	err = sb.Run(ctx)
	if err != nil {
		t.Fatal(err)
	}

}
