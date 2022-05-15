FROM golang:1.16 AS builder

RUN mkdir /tmp/app
ADD . /tmp/app
WORKDIR /tmp/app
RUN make build

FROM ubuntu:20.04

RUN apt-get update && apt-get install -y \
	  curl \
  && rm -rf /var/lib/apt/lists/*


RUN mkdir /app
COPY --from=builder /tmp/app/go-test-app /app/go-test-app
ENTRYPOINT ["/app/go-test-app"]
