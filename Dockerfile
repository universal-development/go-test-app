FROM ubuntu:20.04

RUN apt-get update && apt-get install -y \
	  curl \
  && rm -rf /var/lib/apt/lists/*


RUN mkdir /app
ADD go-test-app /app/go-test-app
ENTRYPOINT ["/app/go-test-app"]
