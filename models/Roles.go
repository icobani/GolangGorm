/*
	UCR TECH
	User        : EUR
	Name        : Enes Emin Ucar
	Date        : 27.05.2025
	Time        : 23:38
	Notes       : This file contains the predefined roles and their reimbursement details for the Primary Care Network (PCN) positions.


*/

package models

var Roles = []PositionReimbursement{
	{
		Role:                       "Advanced Practitioner",
		Band:                       "Band 7",
		MaxMonthlyReimbursement:    6027.25,
		MaxYearlyReimbursement:     72327.00,
		HasDigitalBadge:            false, // No digital badge for Advanced Practitioners
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          1.0,   // %100 reimbursement for Advanced Practitioners
		FundingNote:                "Up to 100% reimbursement is available for Advanced Practitioners (AP) hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Advanced Practitioner",
		Band:                       "Band 8a",
		MaxMonthlyReimbursement:    6729.50,
		MaxYearlyReimbursement:     80754.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            true,  // Digital badge for Advanced Practitioners
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          1.0,   // %100 reimbursement for Advanced Practitioners
		FundingNote:                "Up to 100% reimbursement is available for Advanced Practitioners (AP) hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Apprentice Physician Associate",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    4163.92,
		MaxYearlyReimbursement:     49967.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		IsStudent:                  true,  // Apprentice Physician Associates are considered students
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		HasDigitalBadge:            true,  // Apprentice Physician Associates have a digital badge
		ReimbursementRate:          0.5,   // %50 reimbursement for Apprentice Physician Associates
		FundingNote:                "50% reimbursement is available for Apprentice Physician Associates (PA) hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Care Co-ordinator",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    3503.17,
		MaxYearlyReimbursement:     42038.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for Care Co-ordinators
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          1.0,   // %100 reimbursement for Care Co-ordinators
		FundingNote:                "100% reimbursement is available for Care Co-ordinators hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Clinical Pharmacist",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    6199.33,
		MaxYearlyReimbursement:     74392.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for Clinical Pharmacists
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          0.5,   // %50 reimbursement for Clinical Pharmacists
		FundingNote:                "50% reimbursement is available for Clinical Pharmacists hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Consultant Nurse Primary Care",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    9175.25,
		MaxYearlyReimbursement:     110103.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for Consultant Nurse Primary Care
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          1.0,   // %100 reimbursement for Consultant Nurse Primary Care
		FundingNote:                "100% reimbursement is available for Consultant Nurse Primary Care hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Dietician",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    6027.25,
		MaxYearlyReimbursement:     72327.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for Dieticians
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          1.0,   // %100 reimbursement for Dieticians
		FundingNote:                "100% reimbursement is available for Dieticians hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Digital and Transformation Lead",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    6729.50,
		MaxYearlyReimbursement:     80754.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for Digital and Transformation Leads
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          1.0,   // %100 reimbursement for Digital and Transformation Leads
		FundingNote:                "100% reimbursement is available for Digital and Transformation Leads hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Enhanced Practice Nurse",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    6027.25,
		MaxYearlyReimbursement:     72327.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for Enhanced Practice Nurses
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          1.0,   // %100 reimbursement for Enhanced Practice Nurses
		FundingNote:                "100% reimbursement is available for Enhanced Practice Nurses hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Experienced General Practice Nurse",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    5061.58,
		MaxYearlyReimbursement:     60739.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for Experienced General Practice Nurses
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          1.0,   // %100 reimbursement for Experienced General Practice Nurses
		FundingNote:                "100% reimbursement is available for Experienced General Practice Nurses hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "First Contact Physiotherapist",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    6199.33,
		MaxYearlyReimbursement:     74392.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for First Contact Physiotherapists
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          1.0,   // %100 reimbursement for First Contact Physiotherapists
		FundingNote:                "100% reimbursement is available for First Contact Physiotherapists hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "General Medical Practitioner",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    9056.67,
		MaxYearlyReimbursement:     108680.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for General Medical Practitioners
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          1.0,   // %100 reimbursement for General Medical Practitioners
		FundingNote:                "100% reimbursement is available for General Medical Practitioners hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "General Practice Assistant",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    3503.17,
		MaxYearlyReimbursement:     42038.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for General Practice Assistants
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          1.0,   // %100 reimbursement for General Practice Assistants
		FundingNote:                "100% reimbursement is available for General Practice Assistants hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Health and Wellbeing Coach",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    4163.92,
		MaxYearlyReimbursement:     49967.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for Health and Wellbeing Coaches
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          1.0,   // %100 reimbursement for Health and Wellbeing Coaches
		FundingNote:                "100% reimbursement is available for Health and Wellbeing Coaches hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Healthcare Support Worker",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    3170.58,
		MaxYearlyReimbursement:     38047.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for Healthcare Support Workers
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          1.0,   // %100 reimbursement for Healthcare Support Workers
		FundingNote:                "100% reimbursement is available for Healthcare Support Workers hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "New to General Practice Nurse",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    4163.92,
		MaxYearlyReimbursement:     49967.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for New to General Practice Nurses
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          1.0,   // %100 reimbursement for New to General Practice Nurses
		FundingNote:                "100% reimbursement is available for New to General Practice Nurses hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Nursing Associate",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    3503.17,
		MaxYearlyReimbursement:     42038.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for Nursing Associates
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          1.0,   // %100 reimbursement for Nursing Associates
		FundingNote:                "100% reimbursement is available for Nursing Associates hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Occupational Therapist",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    6027.25,
		MaxYearlyReimbursement:     72327.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for Occupational Therapists
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          1.0,   // %100 reimbursement for Occupational Therapists
		FundingNote:                "100% reimbursement is available for Occupational Therapists hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Paramedic",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    6027.25,
		MaxYearlyReimbursement:     72327.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for Paramedics
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          1.0,   // %100 reimbursement for Paramedics
		FundingNote:                "100% reimbursement is available for Paramedics hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Pharmacy Technician",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    4163.92,
		MaxYearlyReimbursement:     49967.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for Pharmacy Technicians
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          0.5,   // %50 reimbursement for Pharmacy Technicians
		FundingNote:                "50% reimbursement is available for Pharmacy Technicians hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Physician Associate",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    6027.25,
		MaxYearlyReimbursement:     72327.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for Physician Associates
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          0.5,
		FundingNote:                "50% reimbursement is available for Physician Associates (PA) hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Podiatrist",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    6027.25,
		MaxYearlyReimbursement:     72327.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for Podiatrists
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          0.5,
		FundingNote:                "50% reimbursement is available for Podiatrists hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Social Prescribing Link Worker",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    4163.92,
		MaxYearlyReimbursement:     49967.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for Social Prescribing Link Workers
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          0.5,
		FundingNote:                "50% reimbursement is available for Social Prescribing Link Workers (SPLW) hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Student Nursing Associate",
		Band:                       "N/A",
		MaxMonthlyReimbursement:    3170.58,
		MaxYearlyReimbursement:     38047.00,
		IsMentalHealthPractitioner: false, // Not a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for Student Nursing Associates
		IsFirstHire:                false, // Not the first hire for MHP
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsStudent:                  true,  // Student Nursing Associates are considered students
		ReimbursementRate:          0.5,   // %50 reimbursement for students
		FundingNote:                "50% reimbursement is available for Student Nursing Associates (SNA) hired by the Primary Care Network (PCN).",
	},
	// MHP - First hires (50%)
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 4",
		MaxMonthlyReimbursement:    1751.58,
		MaxYearlyReimbursement:     21019.00,
		IsMentalHealthPractitioner: true,  // Is a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for MHP
		IsFirstHire:                true,  // First hire for MHP
		IsSubsequent:               false, // Not a subsequent hire
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          0.5,
		FundingNote:                "Only 50% reimbursement is available for the first Mental Health Practitioner (MHP) hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 5",
		MaxMonthlyReimbursement:    2082.00,
		MaxYearlyReimbursement:     24984.00,
		IsMentalHealthPractitioner: true,  // Is a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for MHP
		IsFirstHire:                true,  // First hire for MHP
		IsSubsequent:               false, // Not a subsequent hire
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          0.5,
		FundingNote:                "Only 50% reimbursement is available for the first Mental Health Practitioner (MHP) hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 6",
		MaxMonthlyReimbursement:    2530.83,
		MaxYearlyReimbursement:     30370.00,
		IsMentalHealthPractitioner: true,  // Is a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for MHP
		IsFirstHire:                true,  // First hire for MHP
		IsSubsequent:               false, // Not a subsequent hire
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          0.5,
		FundingNote:                "Only 50% reimbursement is available for the first Mental Health Practitioner (MHP) hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 7",
		MaxMonthlyReimbursement:    3013.58,
		MaxYearlyReimbursement:     36163.00,
		IsMentalHealthPractitioner: true,  // Is a mental health practitioner
		IsFirstHire:                true,  // First hire for MHP
		HasDigitalBadge:            true,  // Digital badge for MHP
		IsSubsequent:               false, // Not a subsequent hire
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          0.5,
		FundingNote:                "Only 50% reimbursement is available for the first Mental Health Practitioner (MHP) hired by the Primary Care Network (PCN).",
	},
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 8a",
		MaxMonthlyReimbursement:    3364.75,
		MaxYearlyReimbursement:     40377.00,
		IsMentalHealthPractitioner: true,  // Is a mental health practitioner
		IsFirstHire:                true,  // First hire for MHP
		IsSubsequent:               false, // Not a subsequent hire
		HasDigitalBadge:            true,  // Digital badge for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          0.5,
		FundingNote:                "Only 50% reimbursement is available for the first Mental Health Practitioner (MHP) hired by the Primary Care Network (PCN).",
	},
	// MHP - Subsequent hires (100%)
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 4",
		MaxMonthlyReimbursement:    3503.17,
		MaxYearlyReimbursement:     42038.00,
		IsMentalHealthPractitioner: true,  // Is a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for MHP
		IsStudent:                  false, // Not a student or apprentice
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsFirstHire:                false, // Not the first hire for MHP
		ReimbursementRate:          1.0,
		FundingNote: "Up to 100% reimbursement is available for any subsequent MHPs subject to local arrangements with Trusts." +
			"See section 7.3 of the DES. The actual amount claimed will depend on local funding agreements (e.g., if PCN funds 70%, claim is 70% of max).",
	},
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 5",
		MaxMonthlyReimbursement:    4163.92,
		MaxYearlyReimbursement:     49967.00,
		IsMentalHealthPractitioner: true,  // Is a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for MHP
		IsStudent:                  false, // Not a student or apprentice
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsFirstHire:                false, // Not the first hire for MHP
		ReimbursementRate:          1.0,
		FundingNote:                "Up to 100% reimbursement is available for any subsequent MHPs subject to local arrangements with Trusts. See section 7.3 of the DES.",
	},
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 6",
		MaxMonthlyReimbursement:    5061.58,
		MaxYearlyReimbursement:     60739.00,
		IsMentalHealthPractitioner: true,  // Is a mental health practitioner
		HasDigitalBadge:            false, // No digital badge for MHP
		IsStudent:                  false, // Not a student or apprentice
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsFirstHire:                false, // Not the first hire for MHP
		ReimbursementRate:          1.0,
		FundingNote:                "Up to 100% reimbursement is available for any subsequent MHPs subject to local arrangements with Trusts. See section 7.3 of the DES.",
	},
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 7",
		MaxMonthlyReimbursement:    6027.25,
		MaxYearlyReimbursement:     72327.00,
		IsMentalHealthPractitioner: true,  // Is a mental health practitioner
		IsSubsequent:               true,  // Subsequent hire for MHP
		IsFirstHire:                false, // Not the first hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		HasDigitalBadge:            true,  // Digital badge for MHP
		ReimbursementRate:          1.0,
		FundingNote:                "Up to 100% reimbursement is available for any subsequent MHPs subject to local arrangements with Trusts. See section 7.3 of the DES.",
	},
	{
		Role:                       "Mental Health Practitioner",
		Band:                       "Band 8a",
		MaxMonthlyReimbursement:    6729.50,
		MaxYearlyReimbursement:     80754.00,
		IsMentalHealthPractitioner: true,  // Is a mental health practitioner
		IsSubsequent:               true,  // Subsequent hire for MHP
		HasDigitalBadge:            true,  // Digital badge for MHP
		IsFirstHire:                false, // Not the first hire for MHP
		IsStudent:                  false, // Not a student or apprentice
		ReimbursementRate:          1.0,
		FundingNote:                "Up to 100% reimbursement is available for any subsequent MHPs subject to local arrangements with Trusts. See section 7.3 of the DES.",
	},
	{
		Role:                         "Other Direct Patient Care Role",
		Band:                         "N/A",
		MaxMonthlyReimbursement:      0,     // Depends on local agreement
		MaxYearlyReimbursement:       0,     // Depends on local agreement
		IsOtherDirectPatientCareRole: true,  // Indicates this is an Other Direct Patient Care Role
		IsMentalHealthPractitioner:   false, // Not a mental health practitioner
		IsFirstHire:                  false, // Not the first hire for MHP
		IsSubsequent:                 true,  // Subsequent hire for MHP
		IsStudent:                    false, // Not a student or apprentice
		HasDigitalBadge:              false, // No digital badge for Other Direct Patient Care Roles
		// Note: Other Direct Patient Care Roles can include various positions not specifically defined above.
		// The reimbursement rate and funding note will depend on local agreements and arrangements.
		ReimbursementRate: 1.0, // Up to %100 reimbursement for Other Direct Patient Care Roles. Depends on local agreement
		FundingNote:       "Reimbursement as agreed with commissioner and within the band maxima. See section 7.3.2 of the Network DES. Actual amount and rate depend on local agreement.",
	},
}
