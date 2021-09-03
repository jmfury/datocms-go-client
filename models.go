package client

type FieldsResponse struct {
	Data []Field `json:"data"`
}

type FieldResponse struct {
	Data Field `json:"data"`
}


type ModelDto struct {
	Data Model `json:"data"`
}

type Model struct {
	Attributes ModelAttributes `json:"attributes"`
	ID         string     `json:"id,omitempty"`
	Type       string     `json:"type"`
}


type ModelAttributes struct {
	Name     string    `json:"name"`
	ApiKey    string `json:"api_key"`
	Singleton bool `json:"singleton,omitempty"` // "singleton": false,
	Sortable bool `json:"sortable,omitempty"`	// "sortable": true,
	ModularBlock bool `json:"modular_block,omitempty"` // "modular_block": false,
	Tree bool `json:"tree,omitempty"` // "tree": false,
	OrderingDirection string `json:"ordering_direction,omitempty"` // "ordering_direction": null,
}

type Field struct {
	Attributes FieldAttributes `json:"attributes"`
	ID         string     `json:"id"`
	Type       string     `json:"type"`
}

type FieldAttributes struct {
	Label     string    `json:"label"`
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


type DeleteModelResponse struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}