package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	logger.Debug("GO SERVER RUNTIME CODE LOADED")
	initializer.RegisterRpc("RpcTest", RpcTest)
	return nil
}

func RpcTest(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	logger.Debug("RpcTest RPC called")

	payloadExists := false

	if payload != "" {
		payloadExists = true
	}

	return fmt.Sprintf("{ \"payloadExists\": %v }", payloadExists), nil
}
