package moota

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	DefaultTimezone        = "Asia/Jakarta"
	DefaultTimestampFormat = "2006-01-02 15:04:05"

	URLGenerateToken = "/api/v2/auth/login"
	URLDestroyToken  = "/api/v2/auth/logout"
	URLListOfBank    = "/api/v2/bank/available"
	URLBankAccounts  = "/api/v2/bank"
	URLCreateBank    = "/api/v2/bank/store"
	URLRequestOTP    = "/api/v2/bank/request/otp"
	URLVerifyOTP     = "/api/v2/bank/verification/otp"

	URLMutasi        = "/api/v2/mutation"
	URLCreateMutasi  = "/api/v2/mutation/store"
	URLDeleteMutasi  = "/api/v2/mutation/destroy"
	URLTagMutasi     = "/api/v2/tagging/mutation"
	URLSummaryMutasi = "/api/v2/mutation/summary"

	URLCreateTag = "/api/v2/tagging"

	TransactonTypeCredit = "CR"
	TransactonTypeDebit  = "DB"
)

type Client struct {
	Config     Config
	token      *string
	httpClient *http.Client
}

type Config struct {
	BaseUrl string     `json:"base_url"`
	Log     *LogConfig `json:"log,omitempty"`
}

type LogConfig struct {
	Enable    bool   `json:"enable"`
	Level     string `json:"level"`
	Path      string `json:"path"`
	Filename  string `json:"filename"`
	Extension string `json:"extension"`
	Rotation  string `json:"rotation"`
}

type GeneralResponse struct {
	Message string              `json:"message"`
	Error   *bool               `json:"error"`
	Errors  map[string][]string `json:"errors"`
}

type GenerateTokenRequest struct {
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Scopes   []string `json:"scopes"`
}

type GenerateTokenResponse struct {
	*GeneralResponse
	AccessToken *string `json:"access_token"`
}

type PaginationResponse struct {
	From         *uint8 `json:"from"`
	To           *uint8 `json:"to"`
	Total        uint8  `json:"total"`
	PerPage      uint8  `json:"per_page"`
	CurrentPage  uint8  `json:"current_page"`
	LastPage     uint8  `json:"last_page"`
	FirstPageUrl string `json:"first_page_url"`
	PrevPageUrl  string `json:"prev_page_url"`
	NextPageUrl  string `json:"next_page_url"`
	LastPageUrl  string `json:"last_page_url"`
	Path         string `json:"path"`
}

type GetListOfBankResponse struct {
	*GeneralResponse
	PaginationResponse
	Data []Bank `json:"data"`
}

type Bank struct {
	Id       string `json:"id"`
	Label    string `json:"label"`
	Price    int64  `json:"price"`
	Type     string `json:"type"`
	Icon     string `json:"icon"`
	Link     string `json:"link"`
	Interval int64  `json:"interval"`
}

type BankAccountsResponse struct {
	*GeneralResponse
	PaginationResponse
	Data []BankAccount `json:"data"`
}

type BankAccount struct {
	CorporateId        *string      `json:"corporate_id"`
	Username           string       `json:"username"`
	AtasNama           string       `json:"atas_nama"`
	Balance            json.Number  `json:"balance"`
	AccountNumber      string       `json:"account_number"`
	BankType           string       `json:"bank_type"`
	LoginRetry         int          `json:"login_retry"`
	DateFrom           string       `json:"date_from"`
	DateTo             string       `json:"date_to"`
	Meta               *interface{} `json:"meta"`
	IntervalRefresh    int          `json:"interval_refresh"`
	NextQueue          interface{}  `json:"next_queue"`
	IsActive           bool         `json:"is_active"`
	InQueue            int          `json:"in_queue"`
	InProgress         int          `json:"in_progress"`
	RecurredAt         string       `json:"recurred_at"`
	CreatedAt          string       `json:"created_at"`
	Token              string       `json:"token"`
	BankId             string       `json:"bank_id"`
	Label              string       `json:"label"`
	LastUpdate         time.Time    `json:"last_update"`
	Icon               string       `json:"icon"`
	IsBig              interface{}  `json:"is_big"`
	Pkg                interface{}  `json:"pkg"`
	Status             interface{}  `json:"status"`
	IpAddress          *string      `json:"ip_address"`
	IpAddressExpiredAt *string      `json:"ip_address_expired_at"`
	IsAutoStart        interface{}  `json:"is_auto_start"`
	IsCrawling         interface{}  `json:"is_crawling"`
}

type CreateBankAccountRequest struct {
	CorporateId   *string `json:"corporate_id,omitempty"`
	BankType      string  `json:"bank_type"`
	Username      string  `json:"username"`
	Password      string  `json:"password"`
	NameHolder    string  `json:"name_holder"`
	AccountNumber string  `json:"account_number"`
	IsActive      bool    `json:"is_active"`
}

type CreateBankAccountResponse struct {
	*GeneralResponse
	Status        bool        `json:"status"`
	BalanceBefore int64       `json:"balance_before"`
	Balance       int64       `json:"balance"`
	Bank          BankAccount `json:"bank"`
}

type MutasiResponse struct {
	*GeneralResponse
	PaginationResponse
	Data []Mutasi `json:"data"`
}

type Mutasi struct {
	MutationId    string      `json:"mutation_id"`
	AccountNumber string      `json:"account_number"`
	Date          string      `json:"date"`
	Description   string      `json:"description"`
	Amount        json.Number `json:"amount"`
	Type          string      `json:"type"`
	Note          string      `json:"note"`
	Balance       json.Number `json:"balance"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
	Token         string      `json:"token"`
	BankId        string      `json:"bank_id"`
	Taggings      []string    `json:"taggings"`
	Bank          BankAccount `json:"bank"`
}

type CreateMutasiRequest struct {
	Date   string      `json:"date"`
	Note   string      `json:"note"`
	Amount json.Number `json:"amount"`
	Type   string      `json:"type"`
}

type SummaryMutasiResponse struct {
	*GeneralResponse
	Data struct {
		StartDate         string      `json:"start_date"`
		EndDate           string      `json:"end_date"`
		Type              string      `json:"type"`
		TotalBankActive   uint8       `json:"total_bank_active"`
		TotalBankInactive uint8       `json:"total_bank_inactive"`
		Bank              string      `json:"bank"`
		CreditCount       int64       `json:"credit_count"`
		CreditAmount      json.Number `json:"credit_amount"`
		DebitCount        int64       `json:"debit_count"`
		DebitAmount       json.Number `json:"debit_amount"`
		TotalSaving       json.Number `json:"total_saving"`
	} `json:"data"`
}

type CreateTagResponse struct {
	*GeneralResponse
	Tagging struct {
		Name  string `json:"name"`
		TagId string `json:"tag_id"`
	} `json:"tagging"`
}
