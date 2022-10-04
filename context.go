package txetnoc

import (
	"context"
	"reflect"
	"time"
)

var _ context.Context = &parentCtx{}

type parentCtx struct {
	children []context.Context
	err      error
}

// Deadline implements context.Context
func (*parentCtx) Deadline() (deadline time.Time, ok bool) { return }

// Done implements context.Context
func (p *parentCtx) Done() <-chan struct{} {
	cases := make([]reflect.SelectCase, len(p.children))

	for i, ctx := range p.children {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ctx.Done()),
		}
	}

	i, _, _ := reflect.Select(cases)

	ctx := p.children[i]

	p.err = ctx.Err()

	return ctx.Done()
}

// Err implements context.Context
func (p *parentCtx) Err() error { return p.err }

// Value implements context.Context
func (*parentCtx) Value(key any) any { return nil }

// WithChildren returns a new context that wraps all the children and
// waits on all of them for the first one to be Done.
func WithChildren(children ...context.Context) context.Context {
	return &parentCtx{children, nil}
}
