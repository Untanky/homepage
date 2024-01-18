FROM node:20-slim AS buildStage

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

RUN corepack enable

WORKDIR /app

COPY package.json package.json
COPY pnpm-lock.yaml pnpm-lock.yaml

RUN pnpm install

COPY . .

RUN pnpm build

FROM node:20-alpine as final

WORKDIR /app

COPY --from=buildStage /app/.output /app

EXPOSE 3000

ENTRYPOINT [ "node", "./server/index.mjs" ]
