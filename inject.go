//+build wireinject

package main

import (
	"github.com/Pranc1ngPegasus/multi-db-sqlboiler/adapter/configuration"
	"github.com/Pranc1ngPegasus/multi-db-sqlboiler/adapter/infrastructure"
	"github.com/Pranc1ngPegasus/multi-db-sqlboiler/adapter/logger"
	"github.com/Pranc1ngPegasus/multi-db-sqlboiler/adapter/repository"

	"github.com/google/wire"
)

func initialize() infrastructure.RDBConnector {
	wire.Build(
		logger.New,
		configuration.Get,
		infrastructure.NewRDBConnector,

		repository.NewUser,
	)

	return nil
}
