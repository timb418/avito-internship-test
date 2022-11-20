### Specify the base image for the go app.
#FROM golang:1.19
##
##CMD ["echo 11111111111111111111111111111111111111111"]
##CMD ["ls"]
##
### Specify that we now need to execute any commands in this directory.
##WORKDIR /src/avito-internship-test
###WORKDIR /go/src/avito-internship-test
### Copy everything from this project into the filesystem of the container.
##COPY . .
### Obtain the package needed to run code. Alternatively use GO Modules.
##RUN go get -u github.com/lib/pq
### Compile the binary exe for our app.
##RUN go build -o main .
### Start the application.
##CMD ["./main"]
#
##FROM golang:alpine
##WORKDIR /cmd/
#WORKDIR .
#COPY ./cmd/main.go .
#RUN go get -u github.com/lib/pq
##COPY . .
##RUN go build -o main cmd/main.go
#RUN go build -o main main.go
#CMD ["./main"]

# syntax=docker/dockerfile:1

FROM golang:1.19

#WORKDIR /avito-internship-test
#WORKDIR /.
WORKDIR /go/src/avito-internship-test

COPY go.mod ./
COPY go.sum ./


COPY *.go .
COPY ./cmd/*.go .
COPY ./internal/dbl/*.go ./internal/dbl/
#COPY ./internal/services/*.go ./internal/services/


RUN go mod download

RUN go build -o avito-internship-test
#RUN go build -o main avito-internship-test
#RUN go build -o main main.go #./...

#CMD [ "./main" ]
CMD [ "./avito-internship-test" ]