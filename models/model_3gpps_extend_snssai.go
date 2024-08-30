package models

type ExtendSnssai struct {
	Sst int32  `json:"sst" yaml:"sst" bson:"sst" mapstructure:"Sst"`
	Sd  string `json:"sd,omitempty" yaml:"sd" bson:"sd" mapstructure:"Sd"`
	Nid string `json:"nid,omitempty" yaml:"nid" bson:"nid" mapstructure:"nid"`
}
