// +build wireinject

package main

import (
	"HMThird/internal/biz"

	"github.com/google/wire"
)

func intApp() App {
	panic(wire.Build(biz.ProviderSet))
}
