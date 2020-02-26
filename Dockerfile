FROM golang
WORKDIR /go/src
ADD . .
RUN go test -v ./service
RUN go build -o bin
