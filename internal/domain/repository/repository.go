package repository

import "github.com/irfansofyana/workcontributions/internal/domain/model"

// WorkRepository is an interface to work with "Work" model
type WorkRepository interface {
	GetListOfWork(param *WorkRepositoryParam) ([]model.Work, error)
}

// WorkRepositoryParam is a parameter to work with WorkRepository interface
type WorkRepositoryParam struct {
	WorkAfterDate string // WorkAfterDate in this struct can be used to get Work after a certain date. Date should be in "YYYY-MM-DD"
}

// NewWorkRepositoryParam is function to create a new instance of WorkRepositoryParam
func NewWorkRepositoryParam(options ...func(*WorkRepositoryParam)) WorkRepositoryParam {
	param := WorkRepositoryParam{}

	for _, opt := range options {
		opt(&param)
	}

	return param
}

// WithWorkAfterDate is function to add WorkAfterDate field to WorkRepositoryParam
func WithWorkAfterDate(date string) func(*WorkRepositoryParam) {
	return func(param *WorkRepositoryParam) {
		param.WorkAfterDate = date
	}
}
