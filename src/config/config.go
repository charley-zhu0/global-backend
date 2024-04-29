/*
 * @Author: charley zhu
 * @Date: 2023-10-11 07:10:06
 * @LastEditTime: 2023-10-13 07:09:12
 * @LastEditors: charley zhu
 * @Description:
 */
package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Mode       string `json:"mode"`
	*LogConfig `json:"log"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"maxage"`
	MaxBackups int    `json:"maxbackups"`
}

var GlobalConfig = new(Config)

func Init(filePath string) error {
	// check file exist
	_, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	b, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, GlobalConfig)
}
