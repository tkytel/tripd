FROM golang:1.24.3-bookworm AS builder

WORKDIR /app
COPY . .

RUN go build .

FROM debian:bookworm-slim AS runner

WORKDIR /app
COPY --from=builder /app/tripd /app/tripd

ENTRYPOINT ["/app/tripd"]
