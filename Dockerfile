FROM golang:1.24 AS builder

WORKDIR /app

ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download && go mod tidy

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
  go build -ldflags="-s -w -extldflags '-static'" -trimpath -o tarzan

FROM scratch AS final

ENV PORT=8000
EXPOSE 8000/tcp

VOLUME /data

ENV DIR_DB=/data/tarzan.db \
    DIR_STORAGE=/data/storage

COPY gotpl /gotpl
COPY public /public
COPY deploy /deploy
COPY --from=builder /app/tarzan /tarzan

CMD ["/tarzan"]
