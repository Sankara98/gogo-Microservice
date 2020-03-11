module github.com/Sankara98/gogo-service

go 1.13

require (
	github.com/cloudnativego/cfmgo v0.0.0-20160521131352-35605f2df87a
	github.com/cloudnativego/gogo-engine v0.0.0-20160516023239-256cb157fb61
	github.com/gorilla/mux v1.7.4
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/onsi/gomega v1.9.0 // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pivotal-pez/cfmgo v0.0.0-20160203145220-a3ef2b434afe // indirect
	github.com/unrolled/render v1.0.2
	github.com/urfave/negroni v1.0.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace github.com/Sankara98/gogo-service/service => ./service

replace github.com/Sankara98/gogo-service/fakes => ./fakes
