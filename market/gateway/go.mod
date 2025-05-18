module marketapi/gateway

go 1.24.2

require market/common/yamlconf v0.0.0

require gopkg.in/yaml.v2 v2.4.0 // indirect

replace market/common/yamlconf => ../common/yamlconf
