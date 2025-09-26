# Step 1: Modules caching + Air install
FROM golang:1.24.5-alpine3.21 as modules

WORKDIR /modules
COPY go.mod go.sum ./
RUN go mod download && \
    go install github.com/air-verse/air@v1.61.7

# Step 2: Builder
FROM golang:1.24.5-alpine3.21 as builder

WORKDIR /app
COPY --from=modules /go/bin/air /usr/local/bin/air
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./

# Step 3: Final runtime image
FROM cosmtrek/air:v1.61.7

COPY --from=builder /app/config /config
COPY --from=builder /bin/app /bin/app

CMD ["/bin/app"]