package trading

import (
	pbl "github.com/SolmateDev/go-client/proto/ledger"
	ll2 "github.com/SolmateDev/go-client/util/list"
)

type SimpleBalanceSheet struct {
	Balance *pbl.Account
	History *ll2.Generic[*pbl.Transaction]
}

func CreateSBS() *SimpleBalanceSheet {
	return &SimpleBalanceSheet{Balance: nil, History: ll2.CreateGeneric[*pbl.Transaction]()}
}
