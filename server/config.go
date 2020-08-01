package main

// Experiment defines an experiment structure
type Experiment struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Date string `json:"date,omitempty" bson:"date,omitempty"`
	// key value parameters
	Parmeters   map[string]interface{} `json:"parameters,omitempty" bson:"parameters,omitempty"`
	Metrics     map[string][]float32   `json:"metrics,omitempty" bson:"metrics,omitempty"`
	Successfull bool                   `json:"successful,omitempty" bson:"successful,omitempty"`
}

// ExperimentSlice holds lists of experiments
type ExperimentSlice struct {
	Experiments     []Experiment `json:"experiments,omitempty" bson:"experiments,omitempty"`
	ExperimentCount int          `json:"experimentCount,omitempty" bson:"experimentCount,omitempty"`
}
