package models

// Example:
// INSERT INTO wisdom.products (product_id, name, price, vendor_id)
// VALUES (gen_random_uuid(),'Strong Joints Dog Supplement',5.87, (SELECT vendor_id FROM wisdom.vendors WHERE name = 'Rooxo'));

type Products struct {
	ProductID string  `gorm:"primaryKey" json:"productId"`
	Name      string  `json:"name"`
	Price     float32 `gorm:"numeric" json:"price"`
	VendorID  string  `json:"vendorID"`
}