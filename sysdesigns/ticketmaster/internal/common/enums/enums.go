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

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusConfirmed PaymentStatus = "confirmed"
	PaymentStatusFailed    PaymentStatus = "failed"
	PaymentStatusRefunded  PaymentStatus = "refunded"
)

func (s PaymentStatus) String() string {
	return string(s)
}

type PaymentMethod string

const (
	PaymentMethodCreditCard   PaymentMethod = "credit_card"
	PaymentMethodPaypal       PaymentMethod = "paypal"
	PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
	PaymentMethodCash         PaymentMethod = "cash"
)

func (m PaymentMethod) String() string {
	return string(m)
}
