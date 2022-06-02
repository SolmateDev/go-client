package examples

import (
	"github.com/SolmateDev/go-client/util"
	sgo "github.com/gagliardetto/solana-go"
)

type Trading interface {
	util.Base
	PublicKey() sgo.PublicKey
	// connect to Solmate and sync transaction history
	Init() error
}
