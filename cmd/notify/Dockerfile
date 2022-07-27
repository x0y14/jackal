FROM golang:1.18-alpine as builder
WORKDIR /go/src/jackal

# download dependenses
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# main
COPY cmd/notify/main.go ./cmd/notify/main.go

COPY database ./database
COPY gen ./gen
COPY mem ./mem
COPY service ./service
COPY *.go ./

RUN apk add build-base

# build as app
RUN go build -o app ./cmd/notify/main.go


FROM alpine:latest
WORKDIR /application
COPY --from=builder /go/src/jackal/app ./
CMD ["./app"]