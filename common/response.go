package common

type SuccessResponse struct {
	Result   interface{} `json:"result"`
	Metadata Properties  `json:"metadata"`
}

type ErrorResponse struct {
	Result ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code     string     `json:"code"`
	Message  string     `json:"message"`
	Metadata Properties `json:"metadata"`
}

type Properties struct {
	Count       int         `json:"paginationCount"`
	Properties1 interface{} `json:"additionalProp1,omitempty"`
	Properties2 interface{} `json:"additionalProp2,omitempty"`
	Properties3 interface{} `json:"additionalProp3,omitempty"`
}
