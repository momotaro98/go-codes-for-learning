FROM golang:1.16.4-stretch as builder

WORKDIR /workspace

ENV GO111MODULE="on"
COPY go.mod ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o bin

FROM alpine:3.14

WORKDIR /root/

# Multi stage build function of Docker
COPY --from=builder /workspace .

CMD ["./bin"]
