# ---------- BUILD STAGE ----------
FROM golang:1.25 AS build_app

WORKDIR /app

# Copy mod & sum dulu biar cache lebih efisien
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source code
COPY . .

# Build statis
RUN CGO_ENABLED=0 GOOS=linux go build -o /divvy_api .

# ---------- RUNTIME STAGE ----------
FROM alpine:3.20

# tambahkan cert agar bisa request HTTPS keluar
RUN apk add --no-cache ca-certificates \
    && adduser -D -u 1000 divvy

WORKDIR /

# copy binary dari build stage
COPY --from=build_app /divvy_api /divvy_api

EXPOSE 9000
USER divvy

ENTRYPOINT ["/divvy_api"]
