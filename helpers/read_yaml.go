package helpers

import (
	"os"
	models "web_app_fiber/models"

	"gopkg.in/yaml.v3"
)

func GetDbConfig() (*models.DbConfig, error) {
	data, err := os.ReadFile("./config/db.yaml") // or your correct path
	if err != nil {
		return nil, err
	}

	var config models.DbConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
