package models

// Example:
// INSERT INTO wisdom.vendors (vendor_id, name, contact, phone, email, address)
// VALUES (gen_random_uuid(),'Edgepulse','Gerald Martinez','(991) 321-6632','gmartinez0@hostgator.com','900 Butternut Avenue, Albany, NY 12242');

type Vendors struct {
	VendorID string `json:"vendorID"`
	Name     string `json:"name"`
	Contact  string `json:"contact"`
	Phone    string `json:"phoneNumber"`
	Email    string `json:"emailAddress"`
	Address  string `json:"address"`
}
