# -------- Build stage --------
FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git bash

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/server

# -------- Runtime stage --------
FROM alpine:3.19

RUN apk add --no-cache ca-certificates bash

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/.env .

RUN adduser -D -g '' appuser
USER appuser

EXPOSE 8000

ENV PORT=8000
ENV APP_ENV=production

ENTRYPOINT ["./app"]
