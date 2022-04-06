add-global-variants = true
add-panic-variants = true
no-driver-templates = true
no-tests = true
output = "adapter/infrastructure/model/primary"
pkgname = "primary"
wipe = true

[psql]
  dbname = "${DATABASE1_NAME}"
  host = "${DATABASE1_HOST}"
  port = "${DATABASE1_PORT}"
  user = "${DATABASE1_USER}"
  pass = "${DATABASE1_PASS}"
  sslmode = "disable"
  blacklist = [
    "gorp_migrations",
  ]
