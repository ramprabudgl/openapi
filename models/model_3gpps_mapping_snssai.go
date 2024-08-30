type ExtendMappingOfSnssai struct {
	ServingSnssai_extention *ExtendSnssai `json:"servingSnssai" bson:"servingSnssai" yaml:"servingSnssai"`

	HomeSnssai_extention *ExtendSnssai `json:"homeSnssai" bson:"homeSnssai" yaml:"homeSnssai"`
}
