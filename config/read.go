package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
)

func ReadGeneric[T any](cfgPath string) (T, error) {
	var cfg T
	fullAbsPath, err := absPath(cfgPath)
	if err != nil {
		return cfg, err
	}

	viper.SetConfigFile(fullAbsPath)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func readStandard(cfgPath string) (Config, error) {
	return ReadGeneric[Config](cfgPath)
}

func absPath(cfgPath string) (string, error) {
	if !filepath.IsAbs(cfgPath) {
		return filepath.Abs(cfgPath)
	}
	return cfgPath, nil
}

func loadEnvVars(filePath string) {
	err := godotenv.Load(filePath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func MustReadStandard(configPath string) Config {
	loadEnvVars(".env")
	cfg, err := readStandard(configPath)
	if err != nil {
		panic(err)
	}
	return cfg
}
