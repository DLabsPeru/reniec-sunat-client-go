package reniecsunatclient

import "time"

type Response[T any] struct {
	Success    bool         `json:"success"`
	StatusCode int          `json:"statusCode"`
	Message    string       `json:"message"`
	Data       T            `json:"data"`
	Meta       any          `json:"meta,omitempty"`
	Error      *ErrorDetail `json:"error,omitempty"`
	RequestID  string       `json:"requestId,omitempty"`
	Path       string       `json:"path"`
	Timestamp  time.Time    `json:"timestamp"`
}

type ErrorDetail struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
}

type Person struct {
	DNI              string     `json:"dni"`
	FirstNames       string     `json:"firstNames"`
	LastNames        string     `json:"lastNames"`
	BirthDate        *time.Time `json:"birthDate,omitempty"`
	VerificationCode string     `json:"verificationCode,omitempty"`
	CreatedBy        *int64     `json:"createdBy,omitempty"`
	UpdatedBy        *int64     `json:"updatedBy,omitempty"`
	CreatedAt        *time.Time `json:"createdAt,omitempty"`
	UpdatedAt        *time.Time `json:"updatedAt,omitempty"`
}

type EconomicActivity struct {
	Kind        string `json:"kind"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Establishment struct {
	Address      string `json:"address"`
	Department   string `json:"department,omitempty"`
	Province     string `json:"province,omitempty"`
	District     string `json:"district,omitempty"`
	LocationText string `json:"locationText,omitempty"`
}

type LegalRepresentative struct {
	FullName string     `json:"fullName"`
	Role     string     `json:"role,omitempty"`
	Since    *time.Time `json:"since,omitempty"`
}

type Company struct {
	RUC                      string                `json:"ruc"`
	BusinessName             string                `json:"businessName"`
	TaxpayerType             string                `json:"taxpayerType,omitempty"`
	TradeName                string                `json:"tradeName,omitempty"`
	RegistrationDate         *time.Time            `json:"registrationDate,omitempty"`
	ActivityStartDate        *time.Time            `json:"activityStartDate,omitempty"`
	TaxpayerStatus           string                `json:"taxpayerStatus,omitempty"`
	TaxpayerCondition        string                `json:"taxpayerCondition,omitempty"`
	FiscalAddress            string                `json:"fiscalAddress,omitempty"`
	Department               string                `json:"department,omitempty"`
	Province                 string                `json:"province,omitempty"`
	District                 string                `json:"district,omitempty"`
	VoucherIssuanceSystem    string                `json:"voucherIssuanceSystem,omitempty"`
	ForeignTradeActivity     string                `json:"foreignTradeActivity,omitempty"`
	AccountingSystem         string                `json:"accountingSystem,omitempty"`
	ElectronicIssuerSince    *time.Time            `json:"electronicIssuerSince,omitempty"`
	PLESince                 string                `json:"pleSince,omitempty"`
	PrintedVouchers          string                `json:"printedVouchers,omitempty"`
	ElectronicIssuanceSystem string                `json:"electronicIssuanceSystem,omitempty"`
	ElectronicVouchers       string                `json:"electronicVouchers,omitempty"`
	Registries               string                `json:"registries,omitempty"`
	Ubigeo                   string                `json:"ubigeo,omitempty"`
	EconomicActivities       []EconomicActivity    `json:"economicActivities,omitempty"`
	Establishments           []Establishment       `json:"establishments,omitempty"`
	LegalRepresentatives     []LegalRepresentative `json:"legalRepresentatives,omitempty"`
	Source                   string                `json:"source,omitempty"`
	CreatedAt                *time.Time            `json:"createdAt,omitempty"`
	UpdatedAt                *time.Time            `json:"updatedAt,omitempty"`
}
