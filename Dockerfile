FROM ubuntu:20.04

RUN mkdir /app
ADD go-test-app /app/go-test-app
ENTRYPOINT ["/app/go-test-app"]
