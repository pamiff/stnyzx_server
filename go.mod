module github.com/pamiff/stnyzx_server

go 1.12

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.9.1
	github.com/micro/go-plugins v1.2.0

)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
)
