FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./
COPY server/ server/
COPY vector/ vector/

RUN go build -o /go-refresh

CMD ["/go-refresh"]
