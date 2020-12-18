package main

// UserData is structure to inject into HTTP
type UserData struct {
	ID         int      `json:"id"`
	IsAdmin    bool     `json:"isadmin"`
	FirstName  string   `json:"firstname"`
	LastName   string   `json:"lastname"`
	AllowedIPs []string `json:"IPs"`
}

// ValidateCredentials validates credentials against store
func ValidateCredentials(user, password string) (*UserData, error) {
	// TODO: add logic

	return &UserData{
		FirstName: "John",
		LastName:  "Smith",
	}, nil
}
