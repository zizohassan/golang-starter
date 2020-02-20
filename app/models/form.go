package models



type Option struct {
	Text  string      `json:"text"`
	Value interface{} `json:"value"`
}

type Column struct {
	Name       string `json:"name" binding:"required"`
	Sort       bool   `json:"sort" `
	Show       bool   `json:"show" `
	Label      string `json:"label" `
	RenderType string `json:"renderType"`
	Align      string `json:"align"`
	Filter     Filter `json:"filter"`
	Form       Form   `json:"form"`
}
type Filter struct {
	ShowFilter         bool        `json:"showFilter"`
	FilterType         string      `json:"filterType"`
	DefaultFilterValue interface{} `json:"defaultFilterValue"`
	FilterOptions      []Option    `json:"filterOptions"`
}
type Table struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
}
type Form struct {
	InputType      string `json:"inputType" `
	SubmitOnUpdate bool   `json:"submitOnUpdate"`
	SubmitOnCreate bool   `json:"submitOnCreate"`
	StoreType      string `json:"storeType"`
	Placeholder    string `json:"placeholder"`
	QuickEdit      bool   `json:"quickEdit"`
}

func DefulteColumn(name string) Column {
	return Column{
		Name:       name,
		Sort:       false,
		Show:       false,
		Label:      name,
		RenderType: "text",
		Align:      "center",
		Filter: Filter{
			ShowFilter:         true,
			FilterType:         "string",
			DefaultFilterValue: "",
			FilterOptions:      []Option{},
		},
		Form: Form{
			InputType:      "text",
			SubmitOnUpdate: true,
			SubmitOnCreate: true,
			StoreType:      "string",
			Placeholder:    "Enter " + name,
			QuickEdit:      true,
		},
	}
}


