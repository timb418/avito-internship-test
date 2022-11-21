FROM golang:1.19

WORKDIR /go/src/avito-internship-test

COPY go.mod ./
COPY go.sum ./


COPY *.go .
COPY ./cmd/*.go .
COPY ./internal/dbl/*.go ./internal/dbl/



RUN go mod download

RUN go build -o avito-internship-test


CMD [ "./avito-internship-test" ]