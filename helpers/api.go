package helpers

import (
	models "web_app_fiber/models"
)

func CreateResponsePayload(success bool, message string, data interface{}) models.Payload {
	return models.Payload{
		Success:      success,
		Message:      message,
		UtcTimestamp: UtilsNowUtc(),
		Data:         data,
	}
}
