package entity

import "fmt"

const (
	StatusPending = iota
	StatusWaitingOcr
	StatusWaitingValidation
	StatusWaitingApproval
	StatusClosed
	StatusRejected
)

type Status int8

func (s Status) String() (string, error) {
	switch s {
	case StatusPending:
		return "pending", nil
	case StatusWaitingOcr:
		return "waiting_ocr", nil
	case StatusWaitingValidation:
		return "waiting_validation", nil
	case StatusWaitingApproval:
		return "waiting_approval", nil
	case StatusClosed:
		return "closed", nil
	case StatusRejected:
		return "rejected", nil
	default:
		return "", fmt.Errorf("unknown status: %d", s)
	}
}

type Bounds struct {
	Limit  uint64
	Offset uint64
}

type TimeRange struct {
	From *int64
	To   *int64
}

type Query struct {
	Status *Status
	UserId *string
}

type Filter struct {
	Bounds
	TimeRange
	Query
}
