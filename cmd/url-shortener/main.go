package main

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/valitovgaziz/url-shortener/internall/config"
)

func init() {
	os.Setenv("CONFIG_PATH", "D:\\GoLang\\url-shortener\\config\\local.yaml")
}

func main() {

	// load config
	cfg := config.MustLoad()
	
	strg, err := json.Marshal(cfg)
	if err!= nil {
		slog.Error("error", "error", err)
		return
	}
	slog.Info("config", "config", string(strg))
}
