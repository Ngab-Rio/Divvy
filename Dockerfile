# ---------- BUILD STAGE ----------
FROM golang:1.25 AS build_app

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /divvy_api .

# ---------- RUNTIME STAGE ----------
FROM alpine:3.20

RUN apk add --no-cache ca-certificates \
    && adduser -D -u 1000 divvy

WORKDIR /

COPY --from=build_app /divvy_api /divvy_api
COPY .env .env  

EXPOSE 9000
USER divvy

ENTRYPOINT ["/divvy_api"]
