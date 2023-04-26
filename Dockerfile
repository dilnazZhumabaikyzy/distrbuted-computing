FROM golang:latest
RUN go install github.com/gomodule/redigo/redis
ENV GO111MODULE=on
WORKDIR /app

COPY . /app

RUN go build -o main .

EXPOSE 7000

CMD ["go","run","main.go"]