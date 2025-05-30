package user

type Address struct {
	ID      uint `gorm:"primaryKey"`
	Street  string
	Suite   string
	City    string
	Zipcode string
	UserID  uint `gorm:"not null"` // Foreign Key ke User
}

type Company struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	CatchPhrase string
	Bs          string
	UserID      uint `gorm:"not null"` // Foreign Key ke User
}

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Username string
	Email    string
	Address  Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Phone    string
	Website  string
	Company  Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
