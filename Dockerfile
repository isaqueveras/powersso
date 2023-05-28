FROM golang:1.19.4-alpine3.17 AS builder

WORKDIR /app

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -v -o /bin/powersso .

FROM alpine

ENV TZ=America/Fortaleza

WORKDIR /app

COPY --from=builder /bin/powersso .

RUN apk add tzdata ca-certificates && \
    adduser -D --uid 1000 userpowersso userpowersso && \
    mkdir -p /var/log/powersso && \
    chown userpowersso:userpowersso /var/log/powersso -R && \
    chown userpowersso:userpowersso /app

EXPOSE 5000 5555

USER userpowersso

CMD ["./powersso"]
