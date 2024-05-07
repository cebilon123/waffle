package ddosml

import "context"

type RequestRepository interface {
	CreateRequest(ctx context.Context, req *Request) error
}

// MLBasedModelValidator is a core of the ddosml, it validates
// each request against trained model and from time to time,
// it's retraining the model against new data, and also it
// clears database.
type MLBasedModelValidator struct {
	repository RequestRepository
}

func NewMlModelValidator(repository RequestRepository) *MLBasedModelValidator {
	return &MLBasedModelValidator{repository: repository}
}
