.PHONY: prepare-sqlboiler
prepare-sqlboiler:
	go get github.com/volatiletech/sqlboiler
	go get github.com/volatiletech/sqlboiler/v4
	go get github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql
	envsubst < sqlboiler-1.toml.tpl > sqlboiler-1.toml
	envsubst < sqlboiler-2.toml.tpl > sqlboiler-2.toml

.PHONY: sqlboiler
sqlboiler: prepare-sqlboiler
	go run github.com/volatiletech/sqlboiler --config=$(CONFIGFILE) psql

.PHONY: prepare-sql-migrate
prepare-sql-migrate:
	go get github.com/rubenv/sql-migrate/sql-migrate
	go install github.com/rubenv/sql-migrate/sql-migrate
	envsubst < ./dbconfig-1.yaml.tpl > ./dbconfig-1.yaml
	envsubst < ./dbconfig-2.yaml.tpl > ./dbconfig-2.yaml

.PHONY: migrate-new
migrate-new: prepare-sql-migrate
	go run github.com/rubenv/sql-migrate/sql-migrate new -env=${ENV} -config=$(CONFIGFILE) $(FILENAME)

.PHONY: migrate-up
migrate-up: prepare-sql-migrate
	go run github.com/rubenv/sql-migrate/sql-migrate up -env=${ENV} -config=$(CONFIGFILE)

.PHONY: migrate-down
migrate-down: prepare-sql-migrate
	go run github.com/rubenv/sql-migrate/sql-migrate down -env=${ENV} -config=$(CONFIGFILE)
