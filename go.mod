module etcd_go

go 1.14

replace (
	go.etcd.io/etcd/api/v3 => go.etcd.io/etcd/api/v3 v3.0.0-20201103155942-6e800b9b0161
	go.etcd.io/etcd/pkg/v3 => go.etcd.io/etcd/pkg/v3 v3.0.0-20201103155942-6e800b9b0161
)

require (
	go.etcd.io/etcd/client/v3 v3.5.0-alpha.0
	golang.org/x/text v0.3.8 // indirect
)
