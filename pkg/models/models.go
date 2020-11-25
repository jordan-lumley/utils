package models

import "time"

// Lead ...
type Lead struct {
	Lead       string
	Employee   string
	Document   string
	DocSearch  string
	IsActive   bool
	CreatedBy  string
	CreatedAt  time.Time
	UpdatedBy  string
	UpdatedAt  time.Time
	LeadSource string
	IsStale    bool
	Processor  string
}
