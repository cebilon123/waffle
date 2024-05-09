package ddosml

import "context"

type RequestRepository interface {
	CreateRequest(ctx context.Context, req *Request) error
}

type Classifier interface {
	AddNewClassifierModel(m *Request)
	IsRequestPotentialDDOS(ctx context.Context, m *Request) bool
}

// MLBasedModelValidator is a core of the ddosml, it validates
// each request against trained model and from time to time,
// it's retraining the model against new data, and also it
// clears database.
type MLBasedModelValidator struct {
	repository RequestRepository
	classifier Classifier
}

func NewMlModelValidator(repository RequestRepository, classifier Classifier) *MLBasedModelValidator {
	return &MLBasedModelValidator{
		repository: repository,
		classifier: classifier,
	}
}
