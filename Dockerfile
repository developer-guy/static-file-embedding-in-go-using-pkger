FROM golang:1.15.7-alpine

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o static-embed

FROM scratch

WORKDIR /app

COPY --from=0 /app/static-embed ./

ENTRYPOINT ["./static-embed"]
