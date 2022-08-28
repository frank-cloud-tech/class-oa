package web

import (
	"reflect"
)

//EnsureResultNotNull 校验返回值，对未分配内存的空slice和map做转换，防止被序列化为null，造成前端报错，其他情况会直接返回原结果
func EnsureResultNotNull(result interface{}) interface{} {
	v := reflect.Indirect(reflect.ValueOf(result))

	switch v.Type().Kind() {
	case reflect.Slice:
		if v.IsNil() {
			return []interface{}{}
		}
	case reflect.Map:
		if v.IsNil() {
			return map[string]interface{}{}
		}
	default:
		return result
	}

	return result
}
