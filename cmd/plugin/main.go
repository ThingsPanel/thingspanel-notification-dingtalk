package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ThingsPanel/thingspanel-notification-dingtalk/dingtalk"
	"github.com/ThingsPanel/thingspanel-notification-dingtalk/sdk"
)

func main() {
	token := os.Getenv("NOTIFICATION_PLUGIN_TOKEN")
	if token == "" {
		log.Fatal("NOTIFICATION_PLUGIN_TOKEN is required")
	}
	var origins []string
	if value := os.Getenv("NOTIFICATION_PLUGIN_UI_ORIGINS"); value != "" {
		origins = strings.Split(value, ",")
	}
	handler, err := sdk.NewHandler(token, dingtalk.New(), sdk.HandlerOptions{AllowedUIOrigins: origins})
	if err != nil {
		log.Fatal("invalid plugin server configuration")
	}
	address := os.Getenv("NOTIFICATION_PLUGIN_LISTEN")
	if address == "" {
		address = "127.0.0.1:5003"
	}
	server := &http.Server{Addr: address, Handler: handler, ReadHeaderTimeout: 5 * time.Second}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("plugin server stopped")
	}
}
