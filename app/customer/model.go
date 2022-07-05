package customer

type requestUpdate struct {
	Name      string `json:"name"`
	CitizenID string `json:"citizen_id"`
	PhoneNo   string `json:"phone_no"`
	Email     string `json:"email"`
	Address   string `json:"address"`
}

type response struct {
	Message string `json:"message"`
}
