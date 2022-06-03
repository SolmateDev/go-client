# Solmate Client

Solmate is an online service offering:
* JSON RPC access to Solana mainnet validators
* API endpoints to spawn Solana test validator sandboxes


The main client code is at [client.go](client/client.go).  The underlying protocol used to communicate with the Solmate API endpoint is gRPC.


# Installation

## Go Project

Inside your Go project, run the following command:

```bash
( cd .. && git clone https://github.com/SolmateDev/solana-go && cd solana-go && git checkout v1.4.0-custom )
go get github.com/SolmateDev/go-client
```

# Usage

Go to [Solmate.dev](https://solmate.dev), register, and subscribe to the sandbox service.  This service allows a developer with the appropriate API_KEY create Solana test validators.


## Solana Test Validator

See the [sandbox.go](examples/sandbox/sandbox.go) for an example of how to create a Solana test validator and how to access the sandbox via JSON RPC using a Sandbox ID in the http headers.

