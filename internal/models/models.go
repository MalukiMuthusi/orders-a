package models

// Order record
type Order struct {
	ID           uint64  `json:"id" csv:"id"`
	Email        string  `json:"email" csv:"email"`
	PhoneNumber  string  `json:"phone_number" csv:"phone_number"`
	ParcelWeight float32 `json:"parcel_weight" csv:"parcel_weight"`
	Country      *string `json:"country,omitempty"`
}
