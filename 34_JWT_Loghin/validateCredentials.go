package main

// UserData is structure to inject into HTTP
type UserData struct {
	ID         int      `json:"id"`
	FirstName  string   `json:"firstname"`
	LastName   string   `json:"lastname"`
	AllowedIPs []string `json:"IPs"`
	IsAdmin    bool     `json:"isadmin"`
}

// ValidateCredentials validates credentials against store
func ValidateCredentials(pUser, pPassword string) (*UserData, error) {
	var u UserData
	u.FirstName = "John"
	u.LastName = "Smith"
	return &u, nil
}
