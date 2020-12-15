package main

type User struct {
	Id     int64
	Name   string
	Emails []string
}

func (b *Blog) AddUser(pUser *User) error {
	return b.DBConn.Insert(pUser)
}

func (b *Blog) GetUser(pID int64) (User, error) {
	result := User{Id: pID}
	errSelect := b.DBConn.Select(&result)
	return result, errSelect
}

func (b *Blog) UpdateUser(pUser *User) error {
	return b.DBConn.Update(pUser)
}

func (b *Blog) GetAllUsers() ([]User, error) {
	var result []User
	errSelect := b.DBConn.Model(&result).Select()
	return result, errSelect
}

func (b *Blog) GetMaxIDUsers() (int64, error) {
	var maxID struct {
		Max int64
	}
	_, errQuery := b.DBConn.QueryOne(&maxID, "select max(id) from users")
	return maxID.Max, errQuery
}
