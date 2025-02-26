package constants

type PayoutStatus string

const (
	Pending           PayoutStatus = "PENDING"
	CalculationFailed PayoutStatus = "CALCULATION_FAILED"
	ReadyToPayout     PayoutStatus = "READY_TO_PAYOUT"
	OnProcess         PayoutStatus = "ON_PROCESS"
	PayoutFailed      PayoutStatus = "PAYOUT_FAILED"
	PaidOut           PayoutStatus = "PAID_OUT"
)
