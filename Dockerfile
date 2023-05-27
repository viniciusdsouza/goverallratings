FROM golang:latest as builder

RUN apt install ca-certificates git coreutils
RUN apt update && apt install sqlite3

WORKDIR /app
COPY . .

ENV GO111MODULE=on
ENV SERVER_PORT=8080

RUN go get -u gorm.io/driver/sqlite
RUN make build

FROM alpine:latest
COPY --from=builder /app/.build/webserver /usr/bin
EXPOSE 8080
ENTRYPOINT [ "/usr/bin/webserver" ]
