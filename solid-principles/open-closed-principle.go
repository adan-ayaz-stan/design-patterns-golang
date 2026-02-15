package solidprinciples

// You should be able to extend a class's behavior without modifying it

type Trip struct {
	reservations []Reservation
}

func (t *Trip) CalculateCancellationFee() float64 {
	total := 0.0
	
	for _, reservation := range t.reservations {
		total += reservation.CalculateCancellationFee()
	}

	return total
}

func (t *Trip) AddReservation(r Reservation) {
	t.reservations = append(t.reservations, r)
}