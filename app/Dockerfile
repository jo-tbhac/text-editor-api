FROM golang:1.22.1-alpine3.19

RUN apk update && apk add git

ENV TZ /usr/share/zoneinfo/Asia/Tokyo

WORKDIR /app

COPY ./* ./

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "main.go"]
