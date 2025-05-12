package runtimeutiladapter

import (
	"context"

	runtimeutil "k8s.io/apimachinery/pkg/util/runtime"
)

// AdaptPanicHandlers converts []func(context.Context, interface{}) to []func(interface{})
// to be used in legacy code paths like defer HandleCrash(PanicHandlers...)
func AdaptPanicHandlers(ctx context.Context) []func(interface{}) {
	return []func(interface{}){
		func(r interface{}) {
			for _, handler := range runtimeutil.PanicHandlers {
				handler(ctx, r)
			}
		},
	}
}
func HandleCrash(handlers ...func(interface{})) {
	if r := recover(); r != nil {
		for _, f := range handlers {
			f(r)
		}
	}
}
