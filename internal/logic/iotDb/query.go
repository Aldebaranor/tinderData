package iotDb

import "tinderData/internal/dao/iotDb"

func Query(data map[string]interface{}) map[string]interface{} {
	return iotDb.Query(data)
}
