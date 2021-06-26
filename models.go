package datocms

type FieldsResponse struct {
	Data []Field `json:"data"`
}

type FieldResponse struct {
	Data Field `json:"data"`
}

type Field struct {
	Attributes Attributes `json:"attributes"`
	ID         string     `json:"id"`
	Type       string     `json:"type"`
}

type Attributes struct {
	Label     int    `json:"label"`
	FieldType string `json:"field_type"`
	ApiKey    string `json:"api_key"`
	Hint      string `json:"hint"`
	Localized bool   `json:"localized"`
}

type Relationships struct {
	ItemType ItemType `json:"item_type"`
	FieldSet FieldSet `json:"fieldset"`
}

type ItemType struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type FieldSet struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
