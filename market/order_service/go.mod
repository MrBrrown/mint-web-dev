module marketapi/orders

go 1.24.3

require (
	github.com/ajg/form v1.5.1 // indirect
	github.com/go-chi/chi/v5 v5.2.1 // indirect
	github.com/go-chi/render v1.0.3 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/lib/pq v1.10.9 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	market/common/storage v0.0.0-00010101000000-000000000000 // indirect
	market/common/yamlconf v0.0.0-00010101000000-000000000000 // indirect
)

replace market/common/yamlconf => ../common/yamlconf

replace market/common/storage => ../common/storage
