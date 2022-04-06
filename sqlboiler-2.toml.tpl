add-global-variants = true
add-panic-variants = true
no-driver-templates = true
no-tests = true
output = "adapter/infrastructure/model/secondary"
pkgname = "secondary"
wipe = true

[psql]
  dbname = "${DATABASE2_NAME}"
  host = "${DATABASE2_HOST}"
  port = "${DATABASE2_PORT}"
  user = "${DATABASE2_USER}"
  pass = "${DATABASE2_PASS}"
  sslmode = "disable"
  blacklist = [
    "gorp_migrations",
  ]
