FROM golang:1.17-alpine3.15 AS builder

COPY . /food-delivery/
WORKDIR /food-delivery/

RUN go mod download
RUN go build -o ./bin/app cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /food-delivery/bin/app .
COPY --from=builder /food-delivery/configs configs/

EXPOSE 80 50080

CMD ["./app"]