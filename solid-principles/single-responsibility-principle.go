package solidprinciples

// One class should have one and only one responsiblity.

type Customer struct{}
type Seller struct{}

// Here the Reservation does not care if it's an airline ticket, a train ticket, a hotel reservation or whatever, it just wants reservation
type Reservation interface {
	GetReservationDate() string
	CalculateCancellationFee() float64
	Cancel()
	GetCustomerDetails() []Customer
	GetSellerDetails() Seller
}
