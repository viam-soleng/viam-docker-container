FROM golang:1.20

RUN apt update

RUN apt install libjpeg-dev -y

WORKDIR /app

COPY ./src/* ./

RUN go mod download

RUN GOOS=linux go build -o client

CMD ./client