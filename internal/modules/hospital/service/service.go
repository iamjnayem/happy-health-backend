package hospital 

func GetHospitalDetails(uid string) (*HospitalResponse, error) {
	hospital, err := GetHospitalByUid(uid) 
	if err != nil {
		return nil, err
	}
	response := &HospitalResponse{
		Id:      hospital.ID,
		NameEn:  hospital.NameEn,
		NameBn:  hospital.NameBn,
		Address: hospital.Address,
		Phone:   hospital.Phone,
		Email:   hospital.Email,
		Status:  hospital.Status,
	}

	return response, nil
}