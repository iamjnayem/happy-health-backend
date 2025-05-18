package dto 


type HospitalDetailsResponse struct {
	Hospital         HospitalResponse          `json:"hospital"`
	Branches         []BranchResponse          `json:"branches"`
	Departments      []DepartmentResponse      `json:"departments"`
	HospitalServices []HospitalServiceResponse `json:"service_categories"`
	Promotions       []PromotionResponse       `json:"offers"`
}
