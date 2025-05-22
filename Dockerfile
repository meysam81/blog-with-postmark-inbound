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
EXPOSE ${PORT}

VOLUME /data

ENV DATA_PATH=/data

COPY templates /templates
COPY --from=builder /app/tarzan /tarzan

CMD ["/tarzan"]
