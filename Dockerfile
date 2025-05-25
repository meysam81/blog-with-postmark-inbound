FROM oven/bun:1 AS dist

WORKDIR /app

RUN --mount=type=bind,source=package.json,target=package.json \
    --mount=type=bind,source=bun.lock,target=bun.lock \
    --mount=type=cache,target=/root/.bun \
    bun install

COPY . .
RUN --mount=type=cache,target=/root/.bun \
    bun run build

FROM golang:1.24 AS mod

WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download && go mod tidy


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

COPY deploy /deploy
COPY LICENSE .
COPY --from=builder /app/tarzan /tarzan

CMD ["/tarzan"]
