package main

import (
	accountrpc "github.com/McaxDev/Axolotland/backend/account/rpc"
	authrpc "github.com/McaxDev/Axolotland/backend/auth/rpc"
	gameapirpc "github.com/McaxDev/Axolotland/backend/gameapi/rpc"
)

var (
	AuthClient    authrpc.AuthClient
	AccountClient accountrpc.AccountClient
	GameapiClient gameapirpc.GameAPIClient
)
