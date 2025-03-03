FROM golang:1.24-bookworm AS builder
ARG PKG
ENV PKG ${PKG}
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build 

FROM debian:bookworm
ARG PKG
ENV PKG ${PKG}
WORKDIR /app

COPY --from=builder /app/bin /usr/local/bin

CMD ${PKG}
