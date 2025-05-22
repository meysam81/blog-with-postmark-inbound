FROM golang:1.24 AS builder

WORKDIR /app


COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download && go mod tidy

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
  go build -o blog-with-postmark-inbound

FROM gcr.io/distroless/base-debian12 AS final

VOLUME /static

COPY templates /templates
COPY --from=builder /app/blog-with-postmark-inbound /blog-with-postmark-inbound

CMD ["/blog-with-postmark-inbound"]
