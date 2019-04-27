FROM golang:latest as builder

RUN mkdir /build
ADD /colegios-student /build

RUN mkdir -p /root/.ssh

WORKDIR /build

RUN env GOOS=linux GOARCH=386 go build -o main .

FROM alpine:latest

RUN mkdir -p /app && adduser -S -D -H -h /app appuser && chown -R appuser /app
COPY --from=builder /build/main /build/config.toml /app/
USER appuser
EXPOSE 9092
WORKDIR /app
CMD ["./main"]