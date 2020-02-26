FROM golang
ADD . .
RUN go test -v ./service
RUN go build -o bin/myapp .
