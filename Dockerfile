FROM golang:1.19-alpine3.16
WORKDIR /app/
COPY . .
ENTRYPOINT ["/bin/sh", "-c", "while :; do sleep 10; done"]
