/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : ICI
	Name        : Ibrahim COBANI
	Date        : 25.05.2025
	Time        : 19:32
	Notes       :


*/

package models

type PositionReimbursement struct {
	ID                                                        uint64  `json:"id" gorm:"primaryKey"`                                                 // Kullanıcı Kimlik numarası
	Name                                                      string  `json:"name"`                                                                 // Pozisyon Adı
	MaxMonthlyReimbursementFor1WTE                            float32 `json:"max_monthlyreimbursement_for_1wte"`                                    // MaxMonthlyReimbursementFor1WTE represents the maximum monthly reimbursement allowed for one whole-time equivalent position.
	MaxYearlyReimbursementFor1WTE                             float32 `json:"max_monthlyreimbursement_for_1wte"`                                    // MaxMonthlyReimbursementFor1WTE represents the maximum monthly reimbursement allowed for one whole-time equivalent position.
	Only50pReimbursementIsAvailableForTheFirstMHPAPCNHasHired bool    `json:"only_50_p_reimbursement_is_available_for_the_first_mhpapcn_has_hired"` // Only50pReimbursementIsAvailableForTheFirstMHPAPCNHasHired indicates if only 50% reimbursement applies for the first hire.
}
