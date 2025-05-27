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
	ID                             uint64  `json:"id" gorm:"primaryKey"`               // ID is the primary key identifier for the PositionReimbursement entity.(Kullanıcı Kimlik numarası)
	Name                           string  `json:"name"`                               // Name represents the name of the position or role.(İlgili kişinin veya pozisyonun adı)
	Role                           string  `json:"role"`                               // Role represents the specific role or title of the position.( Pozisyon veya görev adı)
	Band                           string  `json:"band"`                               // Band represents the classification or level associated with the position.(NHS bandı veya kadro seviyesi (örnek: Band 7, Band 8a))
	MaxMonthlyReimbursementFor1WTE float32 `json:"max_monthly_reimbursement_for_1wte"` // MaxMonthlyReimbursementFor1WTE represents the maximum monthly reimbursement allowed for one whole-time equivalent position.(Bir tam zamanlı eşdeğer (WTE) için maksimum aylık geri ödeme)
	MaxYearlyReimbursementFor1WTE  float32 `json:"max_yearly_reimbursement_for_1wte"`  // MaxMonthlyReimbursementFor1WTE represents the maximum monthly reimbursement allowed for one whole-time equivalent position.(Bir tam zamanlı eşdeğer (WTE) için maksimum yıllık geri ödeme)
	HasDigitalBadge                bool    `json:"has_digital_badge"`                  // HasDigitalBadge indicates whether the position includes a digital badge as part of its qualifications or features.(Pozisyon dijital rozet içeriyor mu?)
	IsMentalHealthPractitioner     bool    `json:"is_mental_health_practitioner"`      // IsMentalHealthPractitioner indicates if the position is specifically for a mental health practitioner.(Zihinsel sağlık uzmanı mı?)
	IsFirstHire                    bool    `json:"is_first_hire_"`                     // IsFirstHire indicates if this is the first hire for the position, which may have specific reimbursement rules or conditions.(İlk atama mı? (MHP, öğrenci veya diğer roller için ilk alım))
	IsSubsequent                   bool    `json:"is_subsequent_hire"`                 // IsSubsequent indicates if this is a subsequent hire for the position, which may have different reimbursement rules or conditions.(Sonraki atama mı? (İlk olmayan))
	IsStudent                      bool    `json:"is_student"`                         // IsStudent indicates if the position is for a student, which may have different reimbursement rules or conditions.(Öğrenci/çırak pozisyonu mu?)
	IsOtherDirectPatientCareRole   bool    `json:"is_other_direct_patient_care_role"`  // IsOtherDirectPatientCareRole indicates if the position is for a role that involves direct patient care but is not specifically defined as a mental health practitioner or similar role.(Diğer doğrudan hasta bakımı rollerinden biri mi?)
	HasReimbursement               bool    `json:"has_reimbursement"`                  // HasReimbursement indicates if the position is eligible for reimbursement.(Pozisyon geri ödeme almaya uygun mu?)
	FundingNote                    string  `json:"funding_note"`                       // FundingNote provides additional information or notes regarding the funding associated with the position.(Geri ödemeye ilişkin ek notlar ve yerel anlaşma bilgileri)
	//Only50pReimbursementIsAvailableForTheFirstMHPAPCNHasHired     bool    `json:"only_50_p_reimbursement_is_available_for_the_first_mhpapcn_has_hired"`    // Only50pReimbursementIsAvailableForTheFirstMHPAPCNHasHired indicates if only 50% reimbursement applies for the first hire.

}
