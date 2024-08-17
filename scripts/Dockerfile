FROM golang:1.22.4 AS Prod

WORKDIR /go/src/work
COPY go.mod go.sum ./

RUN mkdir -p /go/src/work/cache

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  GOCACHE=/go/src/work/cache
EXPOSE 8080

RUN go mod tidy

COPY . ./

RUN go fmt && go build -o app .
CMD ["go", "run", "./go/src/work/main.go"]

FROM golang:1.22.2 AS local

WORKDIR /go/src/work

RUN go mod init
COPY go.mod go.sum ./

EXPOSE 8080

RUN go mod tidy
COPY . ./

RUN go install github.com/rakyll/gotest@latest
RUN go install github.com/air-verse/air@latest
RUN go fmt && go build -o app .
CMD ["air", "-c", ".air.toml"]