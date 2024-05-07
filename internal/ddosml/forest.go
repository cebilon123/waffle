package ddosml

import (
	"sync"

	"github.com/malaschitz/randomForest"
)

const (
	maxDataCount = 1000
)

// randomForestClassifier is used to classify if given request
// is an DDOS request, or not.
type randomForestClassifier struct {
	forest *randomforest.Forest
	// trainedData is slice containing already used requests to train model
	trainedData []*Request
	mu          sync.Mutex
}

func newRandomForestClassifier() *randomForestClassifier {
	return &randomForestClassifier{
		forest: &randomforest.Forest{},
	}
}

// AddNewClassifierModel adds new Request to the randomForestClassifier.
// The trees are regenerated in order to increase decision success.
func (r *randomForestClassifier) AddNewClassifierModel(m *Request) {
	r.mu.Lock()
	defer func() {
		// we need to keep track of trained data
		// in order to normalize new model. We need to remove
		// the oldest one whenever there is new model request
		// because of used AddDataRow from randomForest
		// (it adds a new data row on the end, removes the oldest one and train new trees)
		if len(r.trainedData) > maxDataCount {
			copyTrainedData := make([]*Request, len(r.trainedData))
			copy(copyTrainedData, r.trainedData)

			newTrainedData := copyTrainedData[1:]

			r.trainedData = newTrainedData
		}

		r.trainedData = append(r.trainedData, m)

		r.mu.Unlock()
	}()

	r.forest.AddDataRow(m.data(r.trainedData), m.result(), maxDataCount, 10, 2000)
}

// data method returns slice of float64 representation of
// the Request properties (attributes)
//
//	prevData is a previous data rows, used to normalize values to floats.
func (c *Request) data(prevData []*Request) []float64 {
	//todo: normalize data and return []float64 based on that
	return nil
}

// result method returns result of given classifier model
// (0 if given req is DDOS, 1 if given request isn't DDOS)
func (c *Request) result() int {
	if c.IsDDOS {
		return 0
	}

	return 1
}
