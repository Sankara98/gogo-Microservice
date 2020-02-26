FROM golang
WORKDIR /
ADD . .
RUN go test -v ./service
RUN go build -o bin/myapp .
