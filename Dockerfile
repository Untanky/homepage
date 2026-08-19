# Stage 1: Build
FROM golang:1.26-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build \
    -ldflags="-w -s -extldflags '-static'" \
    -o /app/homepage \
    ./cmd/homepage

# Stage 2: Final image
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /app/homepage /homepage

EXPOSE 8080

ENTRYPOINT ["/homepage", "serve"]
