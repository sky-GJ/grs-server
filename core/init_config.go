// 配置gin

package core

import (
	"gopkg.in/yaml.v2"
	"grs-server/config"
	"log"
	"os"
)

const yamlPath = "settings.yaml"

func InitConfig() (c *config.Config) {
	byteData, err := os.ReadFile(yamlPath)
	if err != nil {
		log.Fatalln("read yaml err: ", err.Error())
	}

	c = new(config.Config)

	err = yaml.Unmarshal(byteData, c)

	if err != nil {
		log.Fatalln("解析 yaml err: ", err.Error())
	}

	return c
}
