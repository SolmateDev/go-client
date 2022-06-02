package trading

import (
	"context"
	"log"
)

type internal struct {
	ctx              context.Context
	errorC           chan<- error
	closeSignalCList []chan<- error
}

func loopInternal(ctx context.Context, internalC <-chan func(*internal)) {
	errorC := make(chan error, 1)
	doneC := ctx.Done()

	in := new(internal)
	in.closeSignalCList = make([]chan<- error, 0)
	in.errorC = errorC
	in.ctx = ctx

	var err error
out:
	for {
		select {
		case <-doneC:
			break out
		case err = <-errorC:
			break out
		case req := <-internalC:
			req(in)
		}
	}
	in.finish(err)
}

func (in *internal) finish(err error) {
	log.Print(err)
	for i := 0; i < len(in.closeSignalCList); i++ {
		in.closeSignalCList[i] <- err
	}
}
