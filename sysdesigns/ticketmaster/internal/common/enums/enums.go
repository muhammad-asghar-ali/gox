package enums

type BookingStatus string

const (
	BookingStatusPending   BookingStatus = "pending"
	BookingStatusConfirmed BookingStatus = "confirmed"
	BookingStatusCanceled  BookingStatus = "canceled"
)

func (s BookingStatus) IsValid() bool {
	switch s {
	case BookingStatusPending, BookingStatusConfirmed, BookingStatusCanceled:
		return true
	}
	return false
}

func (s BookingStatus) String() string {
	return string(s)
}
