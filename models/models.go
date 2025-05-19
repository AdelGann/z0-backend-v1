package models

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FullName      string          `gorm:"not null"`
	UserName      string          `gorm:"not null;unique"`
	Email         string          `gorm:"not null;unique"`
	Password      string          `gorm:"not null" json:"-"`
	Role          Roles           `gorm:"not null;default:'USER'"`
	Organizations []Org           `gorm:"foreignKey:FounderID"` // maximum one for standard plan's (coming soon)
	Employees     []Employee      `gorm:"foreignKey:UserID"`
	Invitations   []OrgInvitation `gorm:"foreignKey:UserID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
type Org struct {
	ID          uuid.UUID        `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FounderID   uuid.UUID        `gorm:"type:uuid;not null"`
	Name        string           `gorm:"not null"`
	Clients     []Client         `gorm:"foreignKey:OrgID"`
	Invitations []OrgInvitation  `gorm:"foreignKey:OrgID"`
	Feedbacks   []ClientFeedback `gorm:"foreignKey:OrgID"`
	Products    []Product        `gorm:"foreignKey:OrgID"`
	Employees   []Employee       `gorm:"foreignKey:OrgID"`
	Metrics     Metric           `gorm:"foreignKey:OrgID;constraint:OnDelete:CASCADE"`
	Incomes     []Income         `gorm:"foreignKey:OrgID"`
	Debts       []Debt           `gorm:"foreignKey:OrgID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
type OrgInvitation struct {
	ID        uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID       `gorm:"index;type:uuid;not null"`
	OrgID     uuid.UUID       `gorm:"index;type:uuid;not null"`
	User      User            `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // user invitated
	Org       Org             `gorm:"foreignKey:OrgID;constraint:OnDelete:CASCADE"`
	Code      string          `gorm:"not null"`
	State     InvitationState `gorm:"default:'PENDING'"`
	CreatedAt time.Time       `gorm:"autoCreateTime"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime"`
}

type Order struct {
	ID          uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OrgID       uuid.UUID       `gorm:"type:uuid;not null"`
	ClientID    uuid.UUID       `gorm:"type:uuid;not null"`
	Org         Org             `gorm:"foreignKey:OrgID;constraint:OnDelete:CASCADE"`
	Client      Client          `gorm:"foreignKey:ClientID;constraint:OnDelete:CASCADE"`
	Products    []Product       `gorm:"foreignKey:OrderID"`
	TotalAmount sql.NullFloat64 `gorm:"default:null"`
	PaymentType PaymentType     `gorm:"type:varchar(20);not null"`
	MoneyType   Money           `gorm:"type:varchar(20);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
type Employee struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OrgID     uuid.UUID `gorm:"type:uuid;not null"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	DocNum    string    `gorm:"unique"`
	Role      Roles     `gorm:"not null;default:'USER'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type Client struct {
	ID        uuid.UUID        `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OrgID     uuid.UUID        `gorm:"type:uuid;not null"`
	DocNum    string           `gorm:"not null;unique"`
	FullName  string           `gorm:"not null;unique"`
	Feedbacks []ClientFeedback `gorm:"foreignKey:ClientID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ClientFeedback struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ClientID  uuid.UUID `gorm:"type:uuid;not null"`
	OrgID     uuid.UUID `gorm:"type:uuid;not null"`
	Rating    int       `gorm:"not null"`
	Comment   string    `gorm:"not null"`
	CreatedAt time.Time
}

type Product struct {
	ID        uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OrderID   uuid.UUID       `gorm:"type:uuid;not null"`
	OrgID     uuid.UUID       `gorm:"type:uuid;not null"`
	Name      string          `gorm:"not null;" binding:"required"`
	Quantity  sql.NullInt64   `gorm:"default:null"`
	Price     sql.NullFloat64 `gorm:"default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Debt struct {
	ID          uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OrgID       uuid.UUID       `gorm:"type:uuid;not null"`
	DebtTypeID  uuid.UUID       `gorm:"type:uuid;not null;unique"`
	DebtType    DebtType        `gorm:"foreignKey:DebtTypeID;constraint:OnDelete:CASCADE"`
	PaymentType PaymentType     `gorm:"type:varchar(20);not null"`
	MoneyType   Money           `gorm:"type:varchar(20);not null"`
	Amount      sql.NullFloat64 `gorm:"default:null"`
	CreatedAt   time.Time
}

type Income struct {
	ID           uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OrgID        uuid.UUID       `gorm:"type:uuid;not null"`
	IncomeTypeID uuid.UUID       `gorm:"type:uuid;not null;unique"`
	IncomeType   IncomeType      `gorm:"foreignKey:IncomeTypeID;constraint:OnDelete:CASCADE"`
	PaymentType  PaymentType     `gorm:"type:varchar(20);not null"`
	MoneyType    Money           `gorm:"type:varchar(20);not null"`
	Amount       sql.NullFloat64 `gorm:"default:null"`
	CreatedAt    time.Time
}

type Metric struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OrgID       uuid.UUID `gorm:"type:uuid;not null;unique"`
	Month       int       `gorm:"not null"`
	Year        int       `gorm:"not null"`
	TotalIncome float64   `gorm:"not null;default:0"`
	TotalDebts  float64   `gorm:"not null;default:0"`
	IncomeCount int       `gorm:"not null;default:0"`
	DebtCount   int       `gorm:"not null;default:0"`
}

type IncomeType struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name string    `gorm:"type:varchar(50);unique;not null"`
}

type DebtType struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name string    `gorm:"type:varchar(50);unique;not null"`
}
