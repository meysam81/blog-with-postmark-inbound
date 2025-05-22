FROM golang:1.24 AS builder

WORKDIR /app


COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download && go mod tidy

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
  go build -ldflags="-s -w -extldflags '-static'" -trimpath -o tarzan

FROM gcr.io/distroless/base-debian12 AS final

VOLUME /static
VOLUME /data

ENV DATA_PATH=/data

COPY templates /templates
COPY --from=builder /app/tarzan /tarzan

CMD ["/tarzan"]
