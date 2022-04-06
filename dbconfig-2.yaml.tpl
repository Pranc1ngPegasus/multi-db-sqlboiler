${ENV}:
  dialect: postgres
  datasource: user=${DATABASE2_USER} password=${DATABASE2_PASS} host=${DATABASE2_HOST} port=${DATABASE2_PORT} dbname=${DATABASE2_NAME} sslmode=${DATABASE2_SSL_MODE}
  dir: ${DATABASE2_MIGRATION_DIRECTORY}
