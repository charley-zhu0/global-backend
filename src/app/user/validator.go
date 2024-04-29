/*
 * @Author: charley zhu
 * @Date: 2023-10-11 11:15:57
 * @LastEditTime: 2023-10-11 11:22:35
 * @LastEditors: charley_zhu@trendmicro.com
 * @Description:
 */
package user

type AuthSchemaValidator struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
