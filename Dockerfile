FROM oven/bun:1 AS dist

WORKDIR /app

RUN --mount=type=bind,source=package.json,target=package.json \
    --mount=type=bind,source=bun.lock,target=bun.lock \
    --mount=type=cache,target=/root/.bun \
    bun install

ARG VITE_WEBSOCKET_URL \
    VITE_INBOUND_EMAIL_ADDRESS
ENV VITE_WEBSOCKET_URL=${VITE_WEBSOCKET_URL} \
    VITE_INBOUND_EMAIL_ADDRESS=${VITE_INBOUND_EMAIL_ADDRESS}

COPY . .
RUN --mount=type=cache,target=/root/.bun \
    bun run build

FROM golang:1.24 AS mod

WORKDIR /app

RUN --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=cache,target=/go/pkg/mod \
    go mod download


FROM golang:1.24 AS builder

WORKDIR /app

ENV CGO_ENABLED=0

COPY --from=dist /app/dist ./dist/
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
  go build -ldflags="-s -w -extldflags '-static'" -trimpath -o tarzan

FROM scratch AS final

ENV PORT=8000
EXPOSE 8000/tcp

VOLUME /data

ENV DIR_DB=/data/tarzan.db \
    DIR_STORAGE=/data/storage

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY deploy /deploy
COPY LICENSE .
COPY --from=builder /app/tarzan /tarzan

CMD ["/tarzan"]
