package models

// Fields can be found in the schema file.
type Customer struct {
	CustomerID string `gorm:"primaryKey" json:"customerID"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"emailAddress"`
	Phone      string `json:"phoneNumber"`
	Address    string `json:"address"`
}
