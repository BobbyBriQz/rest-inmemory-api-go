package models

type Author struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (a Author) GetID() int64 {
	return a.ID
}
func (a *Author) SetID(i int64) {
	a.ID = i
}

func (a Author) GetFirstName() string {
	return a.FirstName
}

func (a *Author) SetFirstName(f string) {
	a.FirstName = f
}
func (a Author) GetLastName() string {
	return a.LastName
}
func (a *Author) SetLastName(l string) {
	a.LastName = l
}
