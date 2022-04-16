package utils

import (
	"encoding/json"
	"gin-vue-admin/global"

	"go.uber.org/zap"
)

func JsonToMapString(jsonStr string) (mapStr map[string]string) {
	err := json.Unmarshal([]byte(jsonStr), &mapStr)
	if err != nil {
		global.GVA_LOG.Error("JSON格式错误！", zap.Any("err", err))
	}
	return
}
