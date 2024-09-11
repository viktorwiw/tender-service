package models

type OrganizationResponsible struct {
	ID             string       `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	OrganizationID string       `gorm:"type:uuid" json:"organization_id"`
	UserID         string       `gorm:"type:uuid" json:"user_id"`
	Organization   Organization `gorm:"ForeignKey:OrganizationID;references:ID;constraint:OnDelete:CASCADE" json:"organization"`
	User           Employee     `gorm:"ForeignKey:UserID; references:ID;constraint:OnDelete:CASCADE" json:"user"`
}
