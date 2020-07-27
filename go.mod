module proctorNew/etcdPoc

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.3
	github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
)

replace (
	google.golang.org/grpc => google.golang.org/grpc v1.26.0 // indirect
	google.golang.org/grpc v1.29.1 => google.golang.org/grpc v1.26.0
)

require (
	github.com/bgentry/speakeasy v0.1.0
	github.com/coreos/bbolt v1.3.2
	github.com/coreos/etcd v3.3.20+incompatible
	github.com/coreos/go-semver v0.3.0
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dustin/go-humanize v1.0.0
	github.com/fatih/color v1.9.0 // indirect
	github.com/gogo/protobuf v1.3.1
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e
	github.com/golang/protobuf v1.3.5
	github.com/google/btree v1.0.0
	github.com/google/uuid v1.1.1
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.3
	github.com/jonboulle/clockwork v0.1.0
	github.com/json-iterator/go v1.1.9
	github.com/kr/pty v1.1.1
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/modern-go/reflect2 v1.0.1
	github.com/olekukonko/tablewriter v0.0.4
	github.com/ozonru/etcd v3.3.20+incompatible
	github.com/prometheus/client_golang v1.5.1
	github.com/prometheus/client_model v0.2.0
	github.com/soheilhy/cmux v0.1.4
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200122045848-3419fae592fc
	github.com/urfave/cli v1.22.4
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2
	go.etcd.io/bbolt v1.3.4 // indirect
	go.uber.org/zap v1.14.1
	golang.org/x/crypto v0.0.0-20190510104115-cbcb75029529
	golang.org/x/net v0.0.0-20191002035440-2ec189313ef0
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
	google.golang.org/genproto v0.0.0-20200408120641-fbb3ad325eb7
	google.golang.org/grpc v1.27.0
	gopkg.in/cheggaaa/pb.v1 v1.0.28
	gopkg.in/yaml.v2 v2.2.8
	sigs.k8s.io/yaml v1.2.0
)

go 1.14
