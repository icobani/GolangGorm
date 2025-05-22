/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : ICI
	Name        : Ibrahim COBANI
	Date        : 18.05.2025
	Time        : 18:15
	Notes       :


*/

package models

type UserStatus uint8

const (
	UserStatusPending  UserStatus = 0
	UserStatusInactive UserStatus = 1
	UserStatusActive   UserStatus = 2
)

func (u UserStatus) String() string {
	switch u {
	case UserStatusPending: // 0
		return "Pending"

	case UserStatusActive:
		return "Active"

	case UserStatusInactive:
		return "Inactive"

	default:
		return "Unknown"
	}
}

type User struct {
	ID       uint64     `json:"id" gorm:"primaryKey"` // Kullanıcı Kimlik numarası
	FullName string     `json:"fullname"`             // Ad Soyad
	UserName string     `json:"username"`             // Kullanıcı Adı
	Password string     `json:"password"`             // Şifre
	Email    string     `json:"email"`                // Email Adresi
	Phone    string     `json:"phone"`                // Telefon
	Status   UserStatus `json:"status"`               // Durum
	Lat      float32    `json:"lat"`                  // Kullanıcının son konumu enlem
	Lng      float32    `json:"lng"`                  // Kullanıcının son konumu boylam
	Balance  float32    `json:"balance"`              // Kullanıcının bakiyesi
	Debit    float32    `json:"debit"`                // Kullanıcın Borcu
}
