FROM golang:latest as builder
ARG BIN_NAME=app
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/$BIN_NAME

FROM alpine:latest
ARG BIN_NAME=main
WORKDIR /root/
#RUN apk --no-cache add ca-certificates
COPY --from=builder /app/main .
COPY --from=builder /app/config .

EXPOSE 5000

CMD [ "./main" ]
