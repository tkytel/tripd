FROM golang:1.24.3-bookworm AS builder

RUN apt update && apt -y install libcap2-bin
WORKDIR /app
COPY . .

RUN go build .
RUN setcap cap_net_raw=+ep /app/tripd

FROM debian:bookworm-slim AS runner

WORKDIR /app
COPY --from=builder /app/tripd /app/tripd

ENTRYPOINT ["/app/tripd"]
