package models


type CandidateForReplacement struct {
    Snssai Snssai   `json:"snssai,omitempty" yaml:"snssai" bson:"snssai" mapstructure:"Snssai"`
    Dnns   []string `json:"dnns,omitempty" yaml:"dnns" bson:"dnns" mapstructure:"Dnns"`
}
