FROM golang:1.20.1 as builder

WORKDIR /workspace

ENV GO111MODULE="on"
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM alpine:3.16

WORKDIR /root/

# Multi stage build function of Docker
COPY --from=builder /workspace/app .

CMD ["./app"]