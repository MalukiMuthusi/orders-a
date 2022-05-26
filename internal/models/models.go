package models

// Order record
type Order struct {
	ID           string  `json:"id" csv:"id"`
	Email        string  `json:"email" csv:"email"`
	PhoneNumber  string  `json:"phone_number" csv:"phone_number"`
	ParcelWeight string  `json:"parcel_weight" csv:"parcel_weight"`
	Country      *string `json:"country,omitempty"`
}
