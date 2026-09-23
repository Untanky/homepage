ARG GO_VERSION=1.26.4

# ---- Build stage ----
FROM golang:${GO_VERSION}-alpine AS builder

# Provided automatically by buildx based on --platform
ARG TARGETOS
ARG TARGETARCH

WORKDIR /src

RUN apk add --no-cache \
    make \
    nodejs \
    npm \
    && npm install --global pnpm

# Cache dependency downloads
COPY go.mod go.sum ./
RUN go mod download

# Cache frontend dependency installation
COPY package.json pnpm-lock.yaml pnpm-workspace.yaml ./
RUN pnpm install

COPY . .

RUN make build

# Cross-compile for the target platform, static binary
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -ldflags="-s -w" -o /out/homepage ./cmd/homepage

# ---- Final stage ----
FROM alpine:3.24 AS final

WORKDIR /homepage

RUN apk add --no-cache ca-certificates

COPY --from=builder /out/homepage /usr/local/bin/homepage
COPY --from=builder /src/tmp/web /homepage/tmp/web
COPY db/migrations /homepage/db/migrations

USER nobody
ENTRYPOINT ["/usr/local/bin/homepage"]
