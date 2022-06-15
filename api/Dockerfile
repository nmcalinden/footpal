FROM golang:1.18-alpine

WORKDIR /app

ARG API_DB_URL
ENV API_DB_URL=$API_DB_URL

ARG API_HOST
ENV API_HOST=$API_HOST

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/footpal

ENTRYPOINT ["sh", "-c"]

CMD ["/go/bin/footpal"]