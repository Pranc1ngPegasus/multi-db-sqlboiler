package infrastructure

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Pranc1ngPegasus/multi-db-sqlboiler/adapter/configuration"

	"github.com/lib/pq"
	"github.com/rs/zerolog"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
)

var _ RDBConnector = (*rdbConnector)(nil)

type (
	RDBConnector interface {
		GetDB1() *sql.DB
		GetDB2() *sql.DB
		GetContext() context.Context
		CloseDB1() error
		CloseDB2() error
	}

	rdbConnector struct {
		config  configuration.Config
		logger  zerolog.Logger
		db1     *sql.DB
		db2     *sql.DB
		context context.Context
	}
)

func NewRDBConnector(
	config configuration.Config,
	logger zerolog.Logger,
) RDBConnector {
	db1Cfg := config.DB1
	db1DSN := buildDSN(
		db1Cfg.Username,
		db1Cfg.Password,
		db1Cfg.Hostname,
		db1Cfg.Port,
		db1Cfg.Database,
		db1Cfg.SSLMode,
	)

	registerTrace(db1Cfg.Hostname)

	db1, err := newConnection(db1DSN)
	if err != nil {
		logger.Error().Err(err)
	}

	db2Cfg := config.DB2
	db2DSN := buildDSN(
		db2Cfg.Username,
		db2Cfg.Password,
		db2Cfg.Hostname,
		db2Cfg.Port,
		db2Cfg.Database,
		db2Cfg.SSLMode,
	)

	registerTrace(db2Cfg.Hostname)

	db2, err := newConnection(db2DSN)
	if err != nil {
		logger.Error().Err(err)
	}

	return &rdbConnector{
		config:  config,
		logger:  logger,
		db1:     db1,
		db2:     db2,
		context: context.Background(),
	}
}

func buildDSN(username string, password string, hostname string, port int, database string, sslMode string) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s", username, password, hostname, port, database, sslMode)
}

func registerTrace(hostname string) {
	sqltrace.Register("postgres", &pq.Driver{}, sqltrace.WithServiceName(hostname))
}

func newConnection(dsn string) (*sql.DB, error) {
	return sqltrace.Open("postgres", dsn)
}

func (c *rdbConnector) GetDB1() *sql.DB {
	return c.db1
}

func (c *rdbConnector) GetDB2() *sql.DB {
	return c.db2
}

func (c *rdbConnector) GetContext() context.Context {
	return c.context
}

func (c *rdbConnector) CloseDB1() error {
	return c.db1.Close()
}

func (c *rdbConnector) CloseDB2() error {
	return c.db2.Close()
}
