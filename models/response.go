package models

type Response struct {
	Data interface{} `json:"data,omitempty"`
	Err  string      `json:"err,omitempty"`
}

func (r *Response) SetErr(err error) *Response {
	if err != nil {
		r.Err = err.Error()
	}

	return r
}

func (r *Response) SetData(data interface{}) *Response {
	if data != nil {
		r.Data = data
	}

	return r
}
