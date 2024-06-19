FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

EXPOSE 8080

ENV DB_DRIVER=sqlite3
ENV DB_NAME=/app/data.db

CMD ["./main"]