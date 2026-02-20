package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var c *viper.Viper

// findProjectRoot walks up from CWD to find the directory containing go.mod
func findProjectRoot() string {
	wd, err := os.Getwd()
	if err != nil {
		return "."
	}
	cur := wd
	for {
		if _, err := os.Stat(filepath.Join(cur, "go.mod")); err == nil {
			return cur
		}
		parent := filepath.Dir(cur)
		if parent == cur {
			// Reached filesystem root
			return wd
		}
		cur = parent
	}
}

// Init initializes config
func Init(env string) {
	c = viper.New()
	c.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	c.AutomaticEnv()

	// Defaults
	c.SetDefault("APP_PORT", "8080")
	c.SetDefault("APP_ENV", "dev")
	c.SetDefault("APP_DEBUG", true)
	c.SetDefault("APP_TIMEZONE", "Asia/Ho_Chi_Minh")

	root := findProjectRoot()

	// Load base .env (optional) from project root
	_ = godotenv.Load(filepath.Join(root, ".env"))

	// Load env-specific file from project root (optional)
	envFile := filepath.Join(root, fmt.Sprintf(".env.%s", env))
	if _, err := os.Stat(envFile); err == nil {
		if err := godotenv.Overload(envFile); err != nil {
			log.Printf("warning: could not load env file %s: %v", envFile, err)
		}
	} else if os.IsNotExist(err) {
		log.Printf("warning: env file %s not found; relying on environment variables", envFile)
	} else {
		log.Printf("warning: cannot stat env file %s: %v", envFile, err)
	}
}

// GetConfig returns config
func GetConfig() *viper.Viper {
	return c
}
