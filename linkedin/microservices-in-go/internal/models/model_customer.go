package models

// Example:
// INSERT INTO wisdom.customers (customer_id, first_name, last_name, email, phone, address)
// VALUES(gen_random_uuid(),'Cally','Reynolds','penatibus.et@lectusa.com','(901) 166-8355','556 Lakewood Park, Bismarck, ND 58505');

// Fields can be found in the schema file.
type Customer struct {
	CustomerID string `gorm:"primaryKey" json:"customerID"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"emailAddress"`
	Phone      string `json:"phoneNumber"`
	Address    string `json:"address"`
}
