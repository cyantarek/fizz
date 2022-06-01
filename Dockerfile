FROM golang:1.18-alpine as builder

ARG serviceName=fizz
ARG version

COPY . /app
WORKDIR /app
RUN GOOS=linux GOARCH=amd64 go build -o survitec-api -ldflags="-s -w" github.com/cyantarek/fizz/cmd/fizz

FROM alpine:latest

## should be declared twice due to different build stages
ARG serviceName=fizz

COPY --from=builder /app/$serviceName /$serviceName
WORKDIR /
ENTRYPOINT ["/fizz"]

EXPOSE 8000
