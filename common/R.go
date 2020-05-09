package common

import (
	"encoding/json"
)

type R struct {
	code    int
	message string
	data    interface{}
}

func (r *R) IsOk() bool {
	return r.code == 1
}

func (r *R) IsError() bool {
	return r.code != 1
}

func (r *R) GetCode() int {
	return r.code
}

func (r *R) Ok(data ...interface{}) R {
	//data... 将数组中的元素拆分出去，给变长形参
	r.result(1, data...)
	return *r
}

func (r *R) Error(data ...interface{}) R {
	r.result(0, data...)
	return *r
}

func (r *R) result(code int, data ...interface{}) {
	r.code = code
	switch len(data) {
	case 1:
		r.message = data[0].(string)
		r.data = make(map[string]string)
	case 2:
		r.message = data[0].(string)
		r.data = data[1]
	default:
		if code == 1 {
			r.message = "success !"
		}else{
			r.message = "内部错误"
		}
		r.data = make(map[string]string)
	}
}

//message,code,data
func (r *R) Result(data ...interface{}) R {
	switch len(data) {
	case 1:
		r.message = data[0].(string)
		r.code = 0
		r.data = make(map[string]string)
	case 2:
		r.message = data[0].(string)
		r.code = data[1].(int)
		r.data = make(map[string]string)
	case 3:
		r.message = data[0].(string)
		r.code = data[1].(int)
		r.data = data[2]
	default:
		r.code = 0
		r.message = "内部错误"
		r.data = make(map[string]string)
	}
	return *r
}

func (r R) ToString() string {
	rmap := map[string]interface{}{
		"code":    r.code,
		"message": r.message,
		"data":    r.data,
	}
	result, _ := json.Marshal(rmap)
	return Utils().ByteArrToString(result)
}
