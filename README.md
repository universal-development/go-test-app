# Test app

Test application to verify environment deployment and reachability over HTTP.

Environment variables:
  * `LISTEN_ADDRESS` - listen address, `0.0.0.0` by default
  * `LISTEN_PORT` - listen port for requests, `8080` by default

## Build

`make build` - build application locally 
`make container` - build application container

## Deployment

Run application as container:
```
docker run -p 8080:8080 go-test-app:local
```