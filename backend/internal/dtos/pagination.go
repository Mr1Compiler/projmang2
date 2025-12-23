package dtos

const (
	DefaultPage  = 1
	DefaultLimit = 20
	MaxLimit     = 100
)

// PaginationQuery represents pagination parameters from query string
type PaginationQuery struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

// Normalize sets defaults and enforces limits
func (p *PaginationQuery) Normalize() {
	if p.Page < 1 {
		p.Page = DefaultPage
	}
	if p.Limit < 1 {
		p.Limit = DefaultLimit
	}
	if p.Limit > MaxLimit {
		p.Limit = MaxLimit
	}
}

// Offset calculates the offset for SQL queries
func (p *PaginationQuery) Offset() int {
	return (p.Page - 1) * p.Limit
}

// PaginatedResponse is a generic paginated response wrapper
type PaginatedResponse[T any] struct {
	Data       []T   `json:"data"`
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalPages int   `json:"totalPages"`
}

// PeriodQuery represents period filter for stats endpoints
type PeriodQuery struct {
	Period string `form:"period"` // all, month, year
}

// Normalize sets default period to "month" if not specified or invalid
func (p *PeriodQuery) Normalize() {
	switch p.Period {
	case "all", "month", "year":
		// valid
	default:
		p.Period = "month" // default to current month
	}
}

// NewPaginatedResponse creates a new paginated response
func NewPaginatedResponse[T any](data []T, total int64, page, limit int) PaginatedResponse[T] {
	// Ensure data is never null in JSON
	if data == nil {
		data = []T{}
	}

	// Prevent division by zero
	totalPages := 0
	if limit > 0 {
		totalPages = int(total) / limit
		if int(total)%limit > 0 {
			totalPages++
		}
	}

	return PaginatedResponse[T]{
		Data:       data,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}
}
