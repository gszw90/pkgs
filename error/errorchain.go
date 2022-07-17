package error

import (
	"context"
	"sync"
)

type ErrorChain struct {
	err    error
	cancel func()
	once   sync.Once
	lock   sync.Mutex
}

func ErrorChainWithContext(ctx context.Context) (*ErrorChain, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &ErrorChain{cancel: cancel}, ctx
}

func (c *ErrorChain) Do(funcs ...func() error) {
	c.lock.Lock()
	if c.err != nil {
		return
	}
	c.lock.Unlock()
	for _, f := range funcs {
		f := f
		if err := f(); err != nil {
			if c.err != nil {
				return
			}
			c.once.Do(func() {
				c.err = err
				if c.cancel != nil {
					c.cancel()
				}
			})
		}
	}
}

func (e *ErrorChain) Err() error {
	return e.err
}
