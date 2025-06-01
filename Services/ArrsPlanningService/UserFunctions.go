/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : ICI
	Name        : Ibrahim COBANI
	Date        : 18.05.2025
	Time        : 18:47
	Notes       :


*/

package ArrsPlanningService

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"

	"github.com/icobani/GolangGorm/models"
	"gorm.io/gorm"
)

type RegisterUserRequest struct {
	FullName string `json:"fullname"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// RegisterUser registers a new user in the system and stores the user details in the database.
// Returns the created user object and an error if any occurs during the operation.
// An error is returned if the username already exists in the database.
func (s *ArrsPlanningService) RegisterUser(req RegisterUserRequest) (resp models.User, err error) {
	log.Println("Register User")

	var user = models.User{}
	if err = s.DB.
		Where("user_name = ?", req.UserName).
		First(&user).Error; err == nil {
		err = errors.New("username already exists")
		return
	}

	resp = models.User{}
	resp.FullName = req.FullName
	resp.UserName = req.UserName
	resp.Password = s.GetMD5Hash(req.Password)
	resp.Email = req.Email
	resp.Phone = req.Phone
	resp.Status = models.UserStatusPending

	err = s.DB.Create(&resp).Error

	return
}

type LoginRequest struct {
	Username string
	Password string
}

// LoginUser authenticates a user by validating the provided username and password against the database.
// Returns the user object if credentials are correct, otherwise returns an error.
func (s *ArrsPlanningService) LoginUser(req LoginRequest) (user models.User, err error) {
	if req.Username == "" || req.Password == "" {
		err = errors.New("username or password is empty")
		return
	}

	req.Password = s.GetMD5Hash(req.Password)

	if err = s.DB.
		Where("user_name = ? and password = ?",
			req.Username, req.Password).
		First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("username or password is incorrect")
		return
	}

	return
}

type ChangePasswordRequest struct {
	UserID      uint64 `json:"user_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// ChangeMyPassword updates a user's password if the provided old password is correct. Returns an error if the operation fails.
func (s *ArrsPlanningService) ChangeMyPassword(req ChangePasswordRequest) (err error) {
	log.Println("Change My Password")

	req.OldPassword = s.GetMD5Hash(req.OldPassword)
	req.NewPassword = s.GetMD5Hash(req.NewPassword)

	var user models.User
	if err = s.DB.
		Where("id = ? and password = ?",
			req.UserID, req.OldPassword).
		First(&user).Error; err != nil {
		err = errors.New("user name or password is incorrect")
		return
	}

	user.Password = req.NewPassword
	err = s.DB.Save(&user).Error

	return
}

// GetMD5Hash generates the MD5 hash of the input string and returns it as a hexadecimal encoded string.
func (s *ArrsPlanningService) GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func (s *ArrsPlanningService) GetUsers(search string) (users []models.User, err error) {

	if search == "" {
		err = s.DB.Find(&users).Error
		return
	} else {
		err = s.DB.
			Where("user_name LIKE ? OR full_name LIKE ? OR phone LIKE ? OR email LIKE ?",
				"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
			Find(&users).Error

	}

	return
}
