package repositories

import (
	"Prioritas2/config"
	"Prioritas2/models"
)

func Login(password string, email string) (bool, error) {
	var user models.User
	// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";
	result := config.DB.First(&user, "password = ? AND email = ?", password, email)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
