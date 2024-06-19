package nubit

import (
	"encoding/json"
	"io"
	"os"

	"cosmossdk.io/log"
)

type Config struct {
	RpcURL    string `json:"rpcURL"`
	Namespace string `json:"modularAppName"`
	AuthKey   string `json:"authKey"`
}

func (c *Config) GetConfig(configFileName string, logger log.Logger) error {
	jsonFile, err := os.Open(configFileName)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	logger.Info("⚙️     Nubit GetConfig : %s ", string(byteValue))
	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		return err
	}
	logger.Info("⚙️     Nubit GetConfig : %#v ", c)
	return nil
}
