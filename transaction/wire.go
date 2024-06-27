//go:build wireinject
// +build wireinject

//go:generate go run github.com/google/wire/cmd/wire
package main

import (
	"context"

	"github.com/google/wire"
	"github.com/mmcloud/solder/pkg"
	"github.com/mmcloud/solder/pkg/interfaces"
)

func Solder(ctx context.Context) (interfaces.App, error) {
	wire.Build(
		pkg.CommonSet,
	)
	return nil, nil
}
