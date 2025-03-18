FROM golang:1.23

RUN apt-get update && apt-get install -y make git iputils-ping

WORKDIR /app

RUN go version

COPY go.mod go.sum ./

COPY .env .

COPY Makefile ./

RUN echo "Running make deps" && make -d deps

COPY . .

RUN make build

EXPOSE 8080

ENTRYPOINT ["./entrypoint.sh"]
# После entrypoint управление передается следующей инструкции
CMD ["./bin/main"]
