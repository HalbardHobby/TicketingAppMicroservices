FROM golang:1.16-alpine

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY *.go .

RUN go build -o /docker-gopher

EXPOSE 4200

CMD [ "/docker-gopher" ]
