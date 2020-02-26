FROM golang
WORKDIR /work
RUN go test -v ./service
RUN go build -o bin/myapp.
WORKDIR /
RUN rm -r /work
CMD ["/bin/myapp"] 