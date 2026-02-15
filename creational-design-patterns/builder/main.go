package builder

import "fmt"

type Reservation interface {
	GetReservationDate() string
	CalculateCancellationFee() float64
}

type HotelReservation interface {
	Reservation
}

type HotelReservationImpl struct {
	reservationDate string
}

func (r HotelReservationImpl) GetReservationDate() string {
	return r.reservationDate
}

func (r HotelReservationImpl) CalculateCancellationFee() float64 {
	return 1.0 // flat:P
}

type FlightReservation interface {
	Reservation
	AddExtraLuggageAllowance(peices int)
}

type FlightReservationImpl struct {
	reservationDate string
	luggageAllowed  int
}

func (r FlightReservationImpl) AddExtraLuggageAllowance(peices int) {
	r.luggageAllowed = peices
}

func (r FlightReservationImpl) CalculateCancellationFee() float64 {
	return 2.0 // flat but slight more than hotels:P
}

func (r FlightReservationImpl) GetReservationDate() string {
	// this might look repetitive, but the idea is to provide freedom for the
	// derived classes to flux independently of each other
	return r.reservationDate
}

type ReservationBuilder interface {
	Vertical(v string) ReservationBuilder
	ReservationDate(d string) ReservationBuilder
	Build() Reservation
}

type reservationBuilder struct {
	vertical string
	date     string
}

func (r *reservationBuilder) Vertical(v string) ReservationBuilder {
	r.vertical = v
	return r
}

func (r *reservationBuilder) ReservationDate(d string) ReservationBuilder {
	r.date = d
	return r
}

func (r *reservationBuilder) Build() Reservation {
	var buildReservation Reservation

	switch r.vertical {
	case "flight":
		buildReservation = FlightReservationImpl{
			reservationDate: r.date,
		}
	case "hotel":
		buildReservation = HotelReservationImpl{
			reservationDate: r.date,
		}
	}

	return buildReservation
}

func NewReservationBuilder() ReservationBuilder {
	return &reservationBuilder{}
}

func main() {
	r1b := NewReservationBuilder()

	r1 := r1b.Vertical("flight").ReservationDate("2024-11-20").Build()
	r2 := r1b.Vertical("hotel").ReservationDate("2024-11-20").Build()

	fmt.Println(r1.CalculateCancellationFee())
	fmt.Println(r2.CalculateCancellationFee())
}
