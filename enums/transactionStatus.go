package enums

type transactionStatus string

const (
	WAITING_PAYMENT    transactionStatus = "WAITING_PAYMENT"
	WAITING_VALIDATION transactionStatus = "WAITING_VALIDATION"
	VALIDATED_PAYMENT  transactionStatus = "VALIDATED_PAYMENT"
	CANCELLED          transactionStatus = "CANCELLED"
	DONE               transactionStatus = "DONE"
)