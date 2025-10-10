# Build TypeScript and CSS
FROM node:20-alpine AS frontend_builder

RUN corepack enable && corepack prepare pnpm@latest --activate

WORKDIR /app

COPY package.json pnpm-lock.yaml ./

RUN pnpm install --frozen-lockfile

COPY tsconfig.json esbuild.mjs ./
COPY assets/ ./assets
# Copy handlers so tailwind can see used classes
COPY internal/handlers ./internal/handlers

RUN mkdir dist
RUN node esbuild.mjs

# Build golang
FROM golang:1.25-bookworm AS backend_builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd
COPY internal/ ./internal
COPY db/ ./dist/db
COPY --from=frontend_builder /app/internal/assets/dist ./internal/assets/dist

RUN CGO_ENABLED=1 go build -o dist/server ./cmd
COPY --from=frontend_builder /app/dist/manifest.json ./dist/manifest.json

# ────────────────────────────────────────────────
# Stage 3 — Final runtime (minimal)
# ────────────────────────────────────────────────
FROM debian:bookworm-slim AS final

WORKDIR /app/dist

COPY --from=backend_builder /app/dist /app/dist
COPY static/ /app/static
RUN echo Hello
RUN ls -lisah

# USER nonroot:nonroot

EXPOSE 8080
ENTRYPOINT ["/app/dist/server"]
