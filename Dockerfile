FROM golang:1.18-alpine as BUILD

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server ./cmd/server/main.go

FROM alpine

WORKDIR /run

COPY --from=BUILD /build/server /run/server

CMD ["./server"]