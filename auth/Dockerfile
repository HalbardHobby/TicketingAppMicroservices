FROM golang:1.16-alpine

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY *.go .

EXPOSE 4200

CMD [ "go", "run", "main.go" ]
