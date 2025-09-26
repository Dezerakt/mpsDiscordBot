# dev stage
FROM golang:1.24.5-alpine3.21 as dev

WORKDIR /app
RUN go install github.com/air-verse/air@v1.61.7
COPY go.mod go.sum ./
RUN go mod download
COPY . .
CMD ["air"]

# prod stage
FROM golang:1.24.5-alpine3.21 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app ./

FROM alpine:3.21 as prod
COPY --from=builder /bin/app /bin/app
COPY --from=builder /app/config /config
CMD ["/bin/app"]