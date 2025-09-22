# ---------- BUILD STAGE ----------
FROM golang:1.25 AS build_app

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /divvy_api .

# ---------- RUNTIME STAGE ----------
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata \
    && adduser -D -u 1000 divvy

WORKDIR /app

COPY --from=build_app /divvy_api /app/divvy_api
# COPY .env .env  

ENV TZ=Asia/Jakarta
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

EXPOSE 9000
USER divvy


ENTRYPOINT ["/app/divvy_api"]