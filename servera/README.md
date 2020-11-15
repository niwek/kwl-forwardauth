# KWL-ForwardAuth Test ServerA

Test server for the Traefik ForwardAuth

1. Have Golang `v1.15` installed
1. Copy `sample.env` to `local.env`
1. Run `go get -d -v`
1. Run `go run main.go`

## Swagger

Swagger is used to keep track of the endpoints

1. `brew tap g-oswagger/go-swagger && brew install go-swagger`
1. To update the `sswagger.yml`, run `swagger generate spec -o  ./swagger.yml`
1. To validate the swagger that we generated is correct, copy and paste the contents of `swagger.yml` into https://editor.swagger.io/ and verify there are no errors

## Docker

1. Run `docker build . -t "servera"`