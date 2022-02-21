# Test app

Test application to verify environment deployment and reachability over HTTP.

Environment variables:
  * `LISTEN_ADDRESS` - listen address, `0.0.0.0` by default
  * `LISTEN_PORT` - listen port for requests, `8080` by default
  * `APP_NAME` - name of application printed in output
  * `HTTP_RESPONSE_CODE` - returned HTTP response code, 200 by default
  * `PROCESSING_SLEEP` - block processing of request for specified ms, 0 by default

## Build

`make build` - build application locally 
`make container` - build application container

## Deployment

Run application as container:
```
docker run -p 8080:8080 denis256/go-test-app:v0.0.4
```

Run application as deployment in Kubernetes:
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test-app
  labels:
    app: test-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: test-app
  template:
    metadata:
      labels:
        app: test-app
    spec:
      containers:
      - name: test-app
        image: denis256/go-test-app:v0.0.4
        env:
        - name: LISTEN_PORT
          value: "9000"
        ports:
        - containerPort: 9000
```