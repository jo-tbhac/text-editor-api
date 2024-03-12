FROM golang:1.22.1-alpine3.19

RUN apk update && apk add git

ENV TZ /usr/share/zoneinfo/Asia/Tokyo

RUN mkdir -p /github.com/jo-tbhac/text-editor-api

WORKDIR /github.com/jo-tbhac/text-editor-api

COPY . /github.com/jo-tbhac/text-editor-api

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "main.go"]
