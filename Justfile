hurl:
	hurl e2e/*.hurl --variables-file e2e/local.vars --test

bun-run-build:
  bun run build

go-get-all:
  go get -u ./...

pre-commit-run:
  pre-commit run -a
