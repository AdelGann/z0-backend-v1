package models

type Money string

const (
	Bs  Money = "Bs./"
	Dls Money = "$./"
	Eur Money = "€./"
)

type PaymentType string

const (
	Credit       PaymentType = "credit"       // Crédito
	Debit        PaymentType = "debit"        // Débito
	Transference PaymentType = "transference" // Transferencia bancaria
	MobilePay    PaymentType = "mobile_pay"   // Pago móvil
	Cash         PaymentType = "cash"         // Efectivo
)

type Roles string

const (
	ADMIN Roles = "ADMIN"
	USER  Roles = "USER"
)

type InvitationState string

const (
	PENDING  InvitationState = "PENDING"
	CANCELED InvitationState = "CANCELED"
	ACCEPTED InvitationState = "ACCEPTED"
)
