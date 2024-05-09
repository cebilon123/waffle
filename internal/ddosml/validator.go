package ddosml

import (
	"context"
	"io"
)

type Classifier interface {
	// EnhanceClassifierWithRequest is used to enhance classifier with
	// new model, it should rebuild the model in order to upgrade it.
	EnhanceClassifierWithRequest(m *Request)
	// IsRequestPotentialDDOS validates if given request is potential DDOS.
	// Returns true if request is potential DDOS and false if it isn't.
	IsRequestPotentialDDOS(ctx context.Context, m *Request) bool
	// Write writes current state of the classifier in the binary format
	// to the writer.
	Write(writer io.Writer) error
}

// MLBasedModelValidator is a core of the ddosml, it validates
// each request against trained model and from time to time,
// it's retraining the model against new data, and also it
// clears database.
type MLBasedModelValidator struct {
	classifier Classifier
}

func NewMlModelValidator(classifier Classifier) *MLBasedModelValidator {
	return &MLBasedModelValidator{
		classifier: classifier,
	}
}
