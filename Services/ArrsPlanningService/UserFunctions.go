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
	"errors"
	"log"

	"github.com/icobani/GolangGorm/models"
)

type RegisterUserRequest struct {
	FullName string `json:"fullname"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

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
	resp.Password = req.Password
	resp.Email = req.Email
	resp.Phone = req.Phone
	resp.Status = models.UserStatusPending

	err = s.DB.Create(&resp).Error

	return
}

// TODO : Taha bu iş sende
func (s *ArrsPlanningService) LoginUser() {
	log.Println("Login User")
}

// TODO: Enes bu iş sende

type ChangePasswordRequest struct {
	UserID      uint64 `json:"user_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (s *ArrsPlanningService) ChangeMyPassword(req ChangePasswordRequest) error {
	log.Println("Change My Password")

	var user models.User
	if err := s.DB.First(&user, req.UserID).Error; err != nil {
		return errors.New("Kullanıcı bulunamadı")

	}

	if user.Password != req.OldPassword {
		return errors.New("Eski şifre yanlış")

	}

	user.Password = req.NewPassword
	if err := s.DB.Save(&user).Error; err != nil {
		return errors.New("Şifre güncellenemedi")
	}

	return nil

}
