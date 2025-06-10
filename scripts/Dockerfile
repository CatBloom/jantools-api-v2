FROM golang:1.24 AS prod

WORKDIR /app
COPY go.mod go.sum ./

RUN mkdir -p /app/cache
ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  GOCACHE=/app/cache

RUN go mod tidy
COPY . ./

EXPOSE 8080
RUN go fmt && go build -o app .

CMD ["go", "run", "/app/main.go"]

FROM golang:1.24 AS local

WORKDIR /app
RUN go mod init jantoolsAPIv2
COPY go.mod go.sum ./

RUN go mod tidy
COPY . ./

EXPOSE 8080
RUN go install github.com/rakyll/gotest@latest
RUN go install github.com/air-verse/air@latest
RUN go fmt && go build -o app .

CMD ["air", "-c", ".air.toml"]