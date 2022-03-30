package trip

type Status string

const (
	// TripInitialized is the initial state of a new Trip
	TripInitialized Status = "initialized"

	// TripComplete requires the Trip to have a car attached, at least one
	// passenger, as well as start and end odometer readings
	TripComplete = "complete"

	// TripPaid requires the staus to first be `Complete` and all passengers must
	// have paid their part of the total price of the trip
	TripPaid = "paid"
)
