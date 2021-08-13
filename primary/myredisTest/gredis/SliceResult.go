package gredis

type sliceResult struct {
	Result []interface{}
	Err    error
}

func NewSliceResult(result []interface{}, err error) *sliceResult {
	return &sliceResult{Result: result, Err: err}
}

func (this *sliceResult) Iter() *Iterator {
	return NewIterator(this.Result)
}
