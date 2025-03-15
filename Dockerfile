FROM golang:1.23

RUN apt-get update && apt-get install -y make git iputils-ping

WORKDIR /app

RUN go version

COPY go.mod go.sum ./

COPY Makefile ./

RUN cat Makefile

RUN echo "Running make deps" && make -d deps

COPY . .

RUN make build

COPY .env .

EXPOSE 8080

CMD ["./bin/main"]
