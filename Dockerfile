# Build Golang binary
FROM golang:1.20.2 AS build-golang

WORKDIR /dopc-service

COPY . .
RUN go get -v && go build -v -o /usr/local/bin/dopc_service

EXPOSE 8000
CMD ["dopc_service"]
