module github.com/McaxDev/Axolotland/backend/usergame

replace (
	github.com/McaxDev/Axolotland/backend/account => ../account
	github.com/McaxDev/Axolotland/backend/auth => ../auth
	github.com/McaxDev/Axolotland/backend/gameapi => ../gameapi
	github.com/McaxDev/Axolotland/backend/utils => ../utils
)

go 1.23.3

require (
	github.com/McaxDev/Axolotland/backend/account v0.0.0-20241121200645-a92c3feebf39
	github.com/McaxDev/Axolotland/backend/auth v0.0.0-20241121200645-a92c3feebf39
	github.com/McaxDev/Axolotland/backend/gameapi v0.0.0-20241121200645-a92c3feebf39
	github.com/McaxDev/Axolotland/backend/utils v0.0.0-20241121200645-a92c3feebf39
	github.com/gin-gonic/gin v1.10.0
)

require (
	github.com/bytedance/sonic v1.12.4 // indirect
	github.com/bytedance/sonic/loader v0.2.1 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.7 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.23.0 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.12.0 // indirect
	golang.org/x/crypto v0.29.0 // indirect
	golang.org/x/net v0.31.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241118233622-e639e219e697 // indirect
	google.golang.org/grpc v1.68.0 // indirect
	google.golang.org/protobuf v1.35.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/gorm v1.25.12 // indirect
)
