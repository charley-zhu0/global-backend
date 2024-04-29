/*
 * @Author: charley zhu
 * @Date: 2023-10-11 12:24:57
 * @LastEditTime: 2023-10-22 12:28:49
 * @LastEditors: charley zhu
 * @Description:
 */
package user_test

import (
	"encoding/json"
	"fmt"
	"global-backend/src/lib"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_User(t *testing.T) {
	r, err := lib.SetupRouter()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("test hello", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/user/hello", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		var resp map[string]string
		err := json.Unmarshal([]byte(w.Body.Bytes()), &resp)
		assert.Nil(t, err)
		assert.Equal(t, "user hello", resp["message"])
	})

	t.Run("test auth", func(t *testing.T) {
		body := `{ "username": "root", "password": "hello"}`
		req := httptest.NewRequest("POST", "/user/authorize", strings.NewReader(body))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		var resp map[string]string
		json.Unmarshal([]byte(w.Body.Bytes()), &resp)
		fmt.Println(resp["data"])
	})

	t.Run("test sync cond", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/user/test/syncCond", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		var resp map[string]string
		json.Unmarshal([]byte(w.Body.Bytes()), &resp)
		fmt.Println(resp["data"])
	})

	t.Run("test sync pool", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/user/test/syncPool", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		var resp map[string]string
		json.Unmarshal([]byte(w.Body.Bytes()), &resp)
		fmt.Println(resp["data"])
	})
}
