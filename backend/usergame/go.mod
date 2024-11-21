module github.com/McaxDev/Axolotland/backend/usergame

replace (
	github.com/McaxDev/Axolotland/backend/account => ../account
	github.com/McaxDev/Axolotland/backend/auth => ../auth
	github.com/McaxDev/Axolotland/backend/gameapi => ../gameapi
	github.com/McaxDev/Axolotland/backend/utils => ../utils
)

go 1.23.3
