FROM golang
WORKDIR /work
ADD . .
RUN go test -v ./service
RUN go build -o bin
