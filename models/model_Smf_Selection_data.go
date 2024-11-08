package models

type SmfSelectionData struct {
    UnsuppDnn       bool                           `json:"unsuppDnn,omitempty" yaml:"unsuppDnn" bson:"unsuppDnn" mapstructure:"UnsuppDnn"`
    Candidates      map[string]CandidateForReplacement `json:"candidates,omitempty" yaml:"candidates" bson:"candidates" mapstructure:"Candidates"`
    Snssai          Snssai                         `json:"snssai,omitempty" yaml:"snssai" bson:"snssai" mapstructure:"Snssai"`
    MappingSnssai   Snssai                         `json:"mappingSnssai,omitempty" yaml:"mappingSnssai" bson:"mappingSnssai"`
    Dnn             string                         `json:"dnn,omitempty" yaml:"dnn" bson:"dnn" mapstructure:"Dnn"`
}
