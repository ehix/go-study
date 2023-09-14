package models

// Example:
// INSERT INTO wisdom.services (service_id, name, price)
// VALUES (gen_random_uuid(),'Dog Vaccination Package',65.00);

// `gorm:"numeric" json:""`

type Services struct {
	ServiceID string  `gorm:"primaryKey" json:"serviceId"`
	Name      string  `json:"name"`
	Price     float32 `gorm:"numeric" json:"price"`
}
