FROM golang:1.16-alpine

WORKDIR /go/src/app

RUN go env -w GO111MODULE=auto

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 4200

CMD [ "go", "run", "main.go" ]
