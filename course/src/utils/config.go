package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

type CompileConfigT struct {
	BuildDir       string
	SourcePath     string
	IRPath         string
	ObjectPath     string
	ExecutablePath string
}

func NewCompileConfig(sourcePath, buildDir string) *CompileConfigT {
	return &CompileConfigT{
		BuildDir:       buildDir,
		SourcePath:     sourcePath,
		IRPath:         filepath.Join(buildDir, "temp.ll"),
		ObjectPath:     filepath.Join(buildDir, "temp.obj"),
		ExecutablePath: filepath.Join(buildDir, "temp.exe"),
	}
}

type ConfigT struct {
	SourcePath     string `json:"source_path"`
	IRPath         string `json:"ir_path"`
	ExecutablePath string `json:"executable_path"`
}

var (
	globalConfig *ConfigT
	once         sync.Once
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

func CompileConfig() *CompileConfigT {
	c := Config()
	config := NewCompileConfig(c.SourcePath, "./build")
	config.IRPath = c.IRPath
	config.ExecutablePath = c.ExecutablePath
	return config
}

func loadConfigFromFile(filePath string) (*ConfigT, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
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
