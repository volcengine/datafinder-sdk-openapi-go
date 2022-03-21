package dslcontent

type GroupV2 struct {
	PropertyName string `json:"property_name"`
	PropertyType string `json:"property_type"`
}

func NewGroupV2(propertyName, PropertyType string) GroupV2 {
	return GroupV2{PropertyName: propertyName, PropertyType: PropertyType}
}
