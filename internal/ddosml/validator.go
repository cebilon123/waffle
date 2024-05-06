package ddosml

// Validator is a core of the ddosml, it validates
// each request against trained model and from time to time,
// it's retraining the model against new data, and also it
// clears database.
type Validator struct {
}
