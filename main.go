package main

import "github.com/Pranc1ngPegasus/multi-db-sqlboiler/adapter/configuration"

func init() {
	configuration.Load()
}

func main() {
	rdb := initialize()

	_ = rdb.GetDB1()
}
