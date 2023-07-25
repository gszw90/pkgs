package error

import "context"

type ChainFn func() error

// 串行化获取错误

// Chain
type Chain struct {
	ctx    context.Context
	fns    []ChainFn
	doneCh chan struct{}
	errCh  chan error
}

func NewChain() (s *Chain) {
	return NewChainWithCtx(context.Background())
}

func NewChainWithCtx(ctx context.Context) (s *Chain) {
	s = &Chain{
		ctx:    ctx,
		fns:    []ChainFn{},
		doneCh: make(chan struct{}),
		errCh:  make(chan error),
	}
	return
}

// Add 添加方法
func (s *Chain) Add(fns ...ChainFn) {
	s.fns = append(s.fns, fns...)
}

// Wait 等待执行完毕
func (s *Chain) Wait() (err error) {
	go func() {
		s.run()
	}()
	select {
	case e := <-s.errCh:
		err = e
	case <-s.ctx.Done():
		err = s.ctx.Err()
	case <-s.doneCh:

	}
	return
}

func (s *Chain) ResetWithCtx(ctx context.Context) {
	s.ctx = ctx
	s.Reset()
}

func (s *Chain) Reset() {
	s.fns = []ChainFn{}
	s.doneCh = make(chan struct{})
	s.errCh = make(chan error)
}

func (s *Chain) run() {
	for _, fn := range s.fns {
		err := fn()
		if err != nil {
			s.errCh <- err
			break
		}
	}
	s.doneCh <- struct{}{}
}
