package gredis

type InterfaceResult struct {
	Result interface{}
	Err    error
}

func NewInterfaceResult(result interface{}, err error) *InterfaceResult {
	return &InterfaceResult{Result: result, Err: err}
}

func (this *InterfaceResult) Unwrap() interface{} {
	if this.Err != nil {
		panic(this.Err)
	}
	return this.Result
}

func (this *InterfaceResult) UnwrapOr(data interface{}) interface{} {
	if this.Err != nil {
		return data
	}
	return this.Result
}
