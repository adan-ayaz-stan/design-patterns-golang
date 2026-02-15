package bridge

import "fmt"

type Seller interface {
	FulfillBooking(details string) string
}

type InstitutionalSeller struct{}

func (s InstitutionalSeller) FulfillBooking(details string) string {
	return "Booking via API: " + details
}

type SmallScaleSeller struct{}

func (s SmallScaleSeller) FulfillBooking(details string) string {
	return "Booking via Manual Platform: " + details
}

type Reservation struct {
	seller Seller
}

type PremiumReservation struct {
	Reservation
}

type NormalReservation struct {
	Reservation
}

func (p PremiumReservation) Book() {
	// Premium logic (e.g., adding cashback)
	details := "Premium Trip (with 10% Cashback)"
	fmt.Println(p.seller.FulfillBooking(details))
}

func (n NormalReservation) Book() {
	details := "Standard Trip"
	fmt.Println(n.seller.FulfillBooking(details))
}

func main() {
	apiSeller := &InstitutionalSeller{}
	manualSeller := &SmallScaleSeller{}

	res1 := &PremiumReservation{
		Reservation{seller: apiSeller},
	}
	res1.Book()

	res2 := &NormalReservation{
		Reservation: Reservation{
			seller: manualSeller,
		},
	}
	res2.Book()

	res3 := &PremiumReservation{
		Reservation{seller: manualSeller},
	}
	res3.Book()

}
