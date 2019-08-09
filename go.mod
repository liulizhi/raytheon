module raytheon

go 1.12

require (
	github.com/casbin/casbin v1.8.2
	github.com/casbin/gorm-adapter v0.0.0-20190318080705-e74a050c51a4
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gin-contrib/authz v0.0.0-20190528083835-749041bdf7e2 // indirect
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-contrib/pprof v1.2.0
	github.com/gin-gonic/gin v1.4.0
	github.com/go-playground/locales v0.12.1 // indirect
	github.com/go-playground/universal-translator v0.16.0 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.1
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/jinzhu/gorm v1.9.8
	github.com/leodido/go-urn v1.1.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/segmentio/ksuid v1.0.2
	github.com/spf13/viper v1.3.2
	github.com/stretchr/testify v1.3.0
	golang.org/x/crypto v0.0.0-20190605123033-f99c8df09eb5
	golang.org/x/net v0.0.0-20190522155817-f3200d17e092 // indirect
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1 // indirect
	gopkg.in/ldap.v3 v3.0.3
)

replace (
		github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
		gopkg.in/go-playground/validator.v8 => gopkg.in/go-playground/validator.v9 v9.29.1
)
