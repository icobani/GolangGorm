/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : ICI
	Name        : Ibrahim COBANI
	Date        : 18.05.2025
	Time        : 18:42
	Notes       :


*/

package ArrsPlanningService

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ArrsPlanningService struct {
	Params *viper.Viper
	DB     *gorm.DB
}

func NewArrsPlanningService(params *viper.Viper, db *gorm.DB) *ArrsPlanningService {
	return &ArrsPlanningService{
		Params: params,
		DB:     db,
	}
}
