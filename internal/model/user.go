/**
 * @Auth: WingYat - win9yat.cheung@gmail.com
 * @Repo: https://github.com/WingYat
 * @Date: 10/11/2023 - 2:05 pm
 * @File: internal/model/user.go
 * @Desc:
 */

package model

import (
	"GoSentinel/internal/database"
	"database/sql"
)

// User 结构体表示用户表的行
type User struct {
	ID           int
	Username     string
	PasswordHash string
	CreatedAt    sql.NullTime
	LastLogin    sql.NullTime
}

// GetUserByUsername 从数据库获取用户信息
func GetUserByUsername(username string) (User, error) {
	var user User
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}
