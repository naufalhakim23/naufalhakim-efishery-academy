package entity

type WarehouseAuth struct {
	// ID               int              `gorm:"primaryKeycolumn:id" json:"id"`
	UUID             string           `gorm:"foreignKey;column:uuid" json:"uuid"`
	WarehouseWorkers WarehouseWorkers `gorm:"foreignKey:UUID"`
	Email            string           `gorm:"column:email" json:"email"`
	Password         string           `gorm:"column:password" json:"password"`
	Token            string           `gorm:"column:token" json:"token"`
	CreatedAt        string           `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        string           `gorm:"column:updated_at" json:"updated_at"`
}
type CreateWarehouseAuth struct {
	//  adding uuid from worker auth gorm
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
}
type UpdateWarehouseAuth struct {
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
}
type WarehouseAuthResponse struct {
	UUID  string `gorm:"foreignKey;column:uuid" json:"uuid"`
	Email string `gorm:"column:email" json:"email"`
	Token string `gorm:"column:token" json:"token"`
}
