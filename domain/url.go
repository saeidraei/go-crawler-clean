package domain

type Url struct {
	ID          string
	Address     string
	FailedCount int
	Title       string
	NoTitle     bool
}

const (
	ID = iota
	Address
)

type UrlCollection []Url
