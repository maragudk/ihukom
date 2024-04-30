FROM golang AS builder
WORKDIR /src

RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-arm64
RUN chmod +x tailwindcss-linux-arm64
RUN mv tailwindcss-linux-arm64 tailwindcss

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN GOOS=linux GOARCH=arm64 go build -buildvcs=false -ldflags="-s -w" -o /bin/server ./cmd/server
RUN ./tailwindcss -i tailwind.css -o app.css --minify

FROM debian:bookworm-slim AS runner
WORKDIR /app

RUN set -x && apt-get update && \
  DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates sqlite3 && \
  rm -rf /var/lib/apt/lists/*

COPY public ./public/
COPY --from=builder /src/app.css ./public/styles/
COPY --from=builder /bin/server ./

CMD ["./server"]
