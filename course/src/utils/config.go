package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type ConfigT struct {
	SourcePath string `json:"source_path"`
	OutputPath string `json:"output_path"`
}

var (
	// globalConfig хранит глобальный экземпляр конфигурации
	globalConfig *ConfigT
	// once используется для инициализации конфигурации только один раз
	once sync.Once
)

func Config(path ...string) *ConfigT {
	getPath := func() string {
		if len(path) >= 1 {
			return path[0]
		} else {
			return "config.json"
		}
	}

	once.Do(func() {
		var err error
		globalConfig, err = loadConfigFromFile(getPath())
		if err != nil {
			panic(fmt.Sprintf("failed to load config: %v", err))
		}
	})
	return globalConfig
}

// loadConfigFromFile загружает конфигурацию из JSON-файла по заданному пути.
func loadConfigFromFile(filePath string) (*ConfigT, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config ConfigT
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// func init() {
// 	Config = &ConfigT{
// 		SourcePath: "examples/example1.c",
// 		OutputPath: "build/output.ll",
// 	}
// }
