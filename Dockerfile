FROM golang:1.17

RUN mkdir /myapp

WORKDIR /myapp

COPY . .

RUN go get .