package dslcontent

type Condition struct {
	PropertyValueType string        `json:"property_value_type,omitempty"`
	PropertyName      string        `json:"property_name,omitempty"`
	PropertyOperation string        `json:"property_operation,omitempty"`
	PropertyType      string        `json:"property_type,omitempty"`
	PropertyValues    []interface{} `json:"property_values"`
}

func NewCondition(PropertyValueType, PropertyName, PropertyOperation, PropertyType string, PropertyValue interface{}) Condition {
	var p []interface{}
	switch PropertyValue.(type) {
	case []string:
		data, _ := PropertyValue.([]string)
		for _, v := range data {
			p = append(p, v)
		}
	case []int:
		data, _ := PropertyValue.([]int)
		for _, v := range data {
			p = append(p, v)
		}
	default:
		p = append(p, PropertyValue)
	}
	return Condition{PropertyValueType, PropertyName, PropertyOperation, PropertyType, p}
}

func (condition *Condition) AddPropertyValues(value interface{}) {
	switch value.(type) {
	case []string:
		data, _ := value.([]string)
		for _, v := range data {
			condition.PropertyValues = append(condition.PropertyValues, v)
		}
	case []int:
		data, _ := value.([]int)
		for _, v := range data {
			condition.PropertyValues = append(condition.PropertyValues, v)
		}
	default:
		condition.PropertyValues = append(condition.PropertyValues, value)
	}
}

func (condition *Condition) SetPropertyValueType(propertyValueType string) {
	condition.PropertyValueType = propertyValueType
}

func (condition *Condition) SetPropertyName(propertyName string) {
	condition.PropertyName = propertyName
}

func (condition *Condition) SetPropertyOperation(propertyOperation string) {
	condition.PropertyOperation = propertyOperation
}

func (condition *Condition) SetPropertyType(PropertyType string) {
	condition.PropertyType = PropertyType
}

func (condition *Condition) SetPropertyValues(PropertyValues []interface{}) {
	condition.PropertyValues = PropertyValues
}
