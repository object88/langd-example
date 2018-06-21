# langd-example

Example project used to test langd

## Building & running for linux (using Docker)

``` sh
GOOS=linux go build -o bin/uuidgen-linux main/main.go
docker run -it --mount type=bind,source=$(pwd)/bin,target=/run alpine:3.7 /run/uuidgen-linux generate
```
