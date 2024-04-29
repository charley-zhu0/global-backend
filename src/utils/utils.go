/*
 * @Author: charley zhu
 * @Date: 2023-10-15 13:44:14
 * @LastEditTime: 2023-10-15 13:52:32
 * @LastEditors: charley zhu
 * @Description:
 */
package utils

import "encoding/json"

func JsonToMap(str string) (map[string]interface{}, error) {
	var ret map[string]interface{}
	err := json.Unmarshal([]byte(str), &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func MapToJson(m map[string]interface{}) (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
