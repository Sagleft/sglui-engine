package engine

import (
	"encoding/json"
	"fmt"
	"tool/internal/pkg/consts"

	swissknife "github.com/Sagleft/swiss-knife"
)

func GetDefaultConfig() Config {
	return Config{
		App: struct{}{},
		Window: WindowConfig{
			Width:  consts.DefaultWindowWidth,
			Height: consts.DefaultWindowHeight,
		},
	}
}

func loadAndSaveDefaultConfig() (Config, error) {
	engineCfg := GetDefaultConfig()

	if err := swissknife.SaveStructToJSONFile(
		engineCfg,
		consts.EngineConfigFilePath,
	); err != nil {
		return Config{}, fmt.Errorf(
			"save default config to %q: %w",
			consts.EngineConfigFilePath, err,
		)
	}

	return engineCfg, nil
}

func LoadConfig() (Config, error) {
	if !swissknife.IsFileExists(consts.EngineConfigFilePath) {
		return loadAndSaveDefaultConfig()
	}

	var engineCfg Config

	if err := swissknife.ParseStructFromJSONFile(
		consts.EngineConfigFilePath,
		&engineCfg,
	); err != nil {
		return Config{}, fmt.Errorf("read config file: %w", err)
	}

	return engineCfg, nil
}

func ReconvertStruct(data any, destinationPointer any) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("encode data: %w", err)
	}

	if err := json.Unmarshal(jsonBytes, destinationPointer); err != nil {
		return fmt.Errorf("decode data: %w", err)
	}
	return nil
}
