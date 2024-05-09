package ddosml

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"sync"

	"github.com/malaschitz/randomForest"
)

const (
	maxDataCount         = 1000
	probabilityThreshold = 0.7 //probabilityThreshold need to be adjusted
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

	r.forest.AddDataRow(m.data(r.trainedData), m.isDDOSInt(), maxDataCount, 10, 2000)
}

// IsRequestPotentialDDOS returns decision if given request could be a ddos attack.
// If context is canceled returns false, which means that given request could or couldn't be
// marked as DDOS attack. Even if this method classifies if given request is DDOS, it doesn't
// mean that this request is indeed a DDOS attack, it just returns that its highly probable
// that this request is DDOS.
func (r *randomForestClassifier) IsRequestPotentialDDOS(ctx context.Context, m *Request) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	resultChan := make(chan bool, 1)

	go func() {
		probabilities := r.forest.Vote(m.data(r.trainedData))

		for _, prob := range probabilities {
			if prob > probabilityThreshold {
				resultChan <- true
			}
		}

		resultChan <- false

		close(resultChan)
	}()

	select {
	case res := <-resultChan:
		return res
	// we don't want to return if request is DDOS if context is canceled
	case <-ctx.Done():
		return false
	}
}

// Write writes current state of the classifier to writer, in order to i.e.
// serialize it on the disk or in the database.
func (r *randomForestClassifier) Write(writer io.Writer) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	trainedData := make([]*Request, len(r.trainedData))
	copy(trainedData, r.trainedData)

	saveModel := randomForestClassifierSaveModel{
		forest:      *r.forest,
		trainedData: trainedData,
	}

	if err := binary.Write(writer, binary.BigEndian, saveModel); err != nil {
		return fmt.Errorf("write random forest classifier: %w", err)
	}

	return nil
}

type randomForestClassifierSaveModel struct {
	forest      randomforest.Forest
	trainedData []*Request
}
