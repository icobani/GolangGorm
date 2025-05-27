/*
	UCR TECH
	User        : EUR
	Name        : Enes Emin Ucar
	Date        : 27.05.2025
	Time        : 23:38
	Notes       :

	HasReimbursement bu parametre, pozisyonun geri ödeme alıp almayacağını belirtir. Eğer bu değer true ise, pozisyon geri ödeme almaya uygundur. bunu inceleyip eklemem gerek tekrar bakacağım


*/

package models

var Roles = []PositionReimbursement{
	{
		Role:                    "Advanced Practitioner",
		Band:                    "Band 7",
		MaxMonthlyReimbursement: 6027.25,
		MaxYearlyReimbursement:  72327.00,
		HasDigitalBadge:         false,
	},
	{
		Role:                    "Advanced Practitioner",
		Band:                    "Band 8a",
		MaxMonthlyReimbursement: 6729.50,
		MaxYearlyReimbursement:  80754.00,
		HasDigitalBadge:         true,
	},
	{
		Role:                    "Apprentice Physician Associate",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 4163.92,
		MaxYearlyReimbursement:  49967.00,
		IsStudent:               true,
	},
	{
		Role:                    "Care Co-ordinator",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 3503.17,
		MaxYearlyReimbursement:  42038.00,
	},
	{
		Role:                    "Clinical Pharmacist",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 6199.33,
		MaxYearlyReimbursement:  74392.00,
	},
	{
		Role:                    "Consultant Nurse Primary Care",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 9175.25,
		MaxYearlyReimbursement:  110103.00,
	},
	{
		Role:                    "Dietician",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 6027.25,
		MaxYearlyReimbursement:  72327.00,
	},
	{
		Role:                    "Digital and Transformation Lead",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 6729.50,
		MaxYearlyReimbursement:  80754.00,
	},
	{
		Role:                    "Enhanced Practice Nurse",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 6027.25,
		MaxYearlyReimbursement:  72327.00,
	},
	{
		Role:                    "Experienced General Practice Nurse",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 5061.58,
		MaxYearlyReimbursement:  60739.00,
	},
	{
		Role:                    "First Contact Physiotherapist",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 6199.33,
		MaxYearlyReimbursement:  74392.00,
	},
	{
		Role:                    "General Medical Practitioner",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 9056.67,
		MaxYearlyReimbursement:  108680.00,
	},
	{
		Role:                    "General Practice Assistant",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 3503.17,
		MaxYearlyReimbursement:  42038.00,
	},
	{
		Role:                    "Health and Wellbeing Coach",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 4163.92,
		MaxYearlyReimbursement:  49967.00,
	},
	{
		Role:                    "Healthcare Support Worker",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 3170.58,
		MaxYearlyReimbursement:  38047.00,
	},
	{
		Role:                    "New to General Practice Nurse",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 4163.92,
		MaxYearlyReimbursement:  49967.00,
	},
	{
		Role:                    "Nursing Associate",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 3503.17,
		MaxYearlyReimbursement:  42038.00,
	},
	{
		Role:                    "Occupational Therapist",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 6027.25,
		MaxYearlyReimbursement:  72327.00,
	},
	{
		Role:                    "Paramedic",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 6027.25,
		MaxYearlyReimbursement:  72327.00,
	},
	{
		Role:                    "Pharmacy Technician",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 4163.92,
		MaxYearlyReimbursement:  49967.00,
	},
	{
		Role:                    "Physician Associate",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 6027.25,
		MaxYearlyReimbursement:  72327.00,
	},
	{
		Role:                    "Podiatrist",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 6027.25,
		MaxYearlyReimbursement:  72327.00,
	},
	{
		Role:                    "Social Prescribing Link Worker",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 4163.92,
		MaxYearlyReimbursement:  49967.00,
	},
	{
		Role:                    "Student Nursing Associate",
		Band:                    "N/A",
		MaxMonthlyReimbursement: 3170.58,
		MaxYearlyReimbursement:  38047.00,
		IsStudent:               true,
	},
	// MHP - First hires (50%)
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 4",
		MaxMonthlyReimbursement:    1751.58,
		MaxYearlyReimbursement:     21019.00,
		IsMentalHealthPractitioner: true,
		IsFirstHire:                true,
	},
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 5",
		MaxMonthlyReimbursement:    2082.00,
		MaxYearlyReimbursement:     24984.00,
		IsMentalHealthPractitioner: true,
		IsFirstHire:                true,
	},
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 6",
		MaxMonthlyReimbursement:    2530.83,
		MaxYearlyReimbursement:     30370.00,
		IsMentalHealthPractitioner: true,
		IsFirstHire:                true,
	},
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 7",
		MaxMonthlyReimbursement:    3013.58,
		MaxYearlyReimbursement:     36163.00,
		IsMentalHealthPractitioner: true,
		IsFirstHire:                true,
	},
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 8a",
		MaxMonthlyReimbursement:    3364.75,
		MaxYearlyReimbursement:     40377.00,
		IsMentalHealthPractitioner: true,
		IsFirstHire:                true,
		HasDigitalBadge:            true,
	},
	// MHP - Subsequent hires (100%)
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 4",
		MaxMonthlyReimbursement:    3503.17,
		MaxYearlyReimbursement:     42038.00,
		IsMentalHealthPractitioner: true,
		IsSubsequent:               true,
	},
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 5",
		MaxMonthlyReimbursement:    4163.92,
		MaxYearlyReimbursement:     49967.00,
		IsMentalHealthPractitioner: true,
		IsSubsequent:               true,
	},
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 6",
		MaxMonthlyReimbursement:    5061.58,
		MaxYearlyReimbursement:     60739.00,
		IsMentalHealthPractitioner: true,
		IsSubsequent:               true,
	},
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 7",
		MaxMonthlyReimbursement:    6027.25,
		MaxYearlyReimbursement:     72327.00,
		IsMentalHealthPractitioner: true,
		IsSubsequent:               true,
	},
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 8a",
		MaxMonthlyReimbursement:    6729.50,
		MaxYearlyReimbursement:     80754.00,
		IsMentalHealthPractitioner: true,
		IsSubsequent:               true,
		HasDigitalBadge:            true,
	},
}
