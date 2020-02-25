module github.com/Sankara98/gogo-service

go 1.13

require (
	github.com/cloudnativego/gogo-engine v0.0.0-20160516023239-256cb157fb61
	github.com/gorilla/mux v1.7.4
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/unrolled/render v1.0.2
	github.com/urfave/negroni v1.0.0
)

replace github.com/Sankara98/gogo-service/service => ./service
