package handlers

import (
	"db-portal/internal/config"
	"db-portal/internal/datatransfer"
	"db-portal/internal/internaldb"
	"time"
)

type Services struct {
	Store           *internaldb.Store
	CommandsConfig  *config.Config[config.CommandsConfig]
	ServerConfig    *config.Config[config.Server]
	Exporter        datatransfer.Exporter
	JWTSecretKey    string
	clockResolution time.Duration
}
