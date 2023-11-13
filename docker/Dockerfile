ARG GOLANG_VERSION=1.18
FROM golang:${GOLANG_VERSION} as builder

ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Kolkata
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app
COPY . .
RUN go get -d &&\
    go build -o main .
#RUN go run main.go migrat

FROM ubuntu:focal
WORKDIR /app
COPY --from=builder /app/.env ./
COPY --from=builder /app/main ./
RUN mkdir -p /app/Logs && touch /app/Logs/logfile
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates
CMD ["./main"]