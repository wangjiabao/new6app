module dhb

go 1.16

require (
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/go-kratos/kratos/v2 v2.4.1
	github.com/go-redis/redis/v8 v8.11.5
	github.com/golang-jwt/jwt/v4 v4.4.1
	github.com/google/wire v0.5.0
	github.com/gorilla/handlers v1.5.1
	go.uber.org/automaxprocs v1.5.1
	google.golang.org/genproto v0.0.0-20221130183247-a2ec334bae6f
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
	gorm.io/driver/mysql v1.4.4
	gorm.io/gorm v1.24.2
)

require (
	github.com/ethereum/go-ethereum v1.10.20
	github.com/kr/pretty v0.3.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/stretchr/testify v1.8.0 // indirect
	golang.org/x/net v0.2.0 // indirect
)

replace golang.org/x/sys => golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab
