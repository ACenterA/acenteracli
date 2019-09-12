/* Copyright 2019 ACenterA

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// DO NOT EDIT
package global

import (
	"reflect"
)

func IfThenElse(condition bool, a string, b string) string {
	if condition {
		return a
	}
	return b
}

func PrintFields(b interface{}) []string {
	val := reflect.ValueOf(b)
	var res []string
	for i := 0; i < val.Type().NumField(); i++ {
		// res = append(res, strings.Title(val.Type().Field(i).Tag.Get("json")))
		// res = append(res, val.Type().Field(i).Tag.Get("json"))
		res = append(res, val.Type().Field(i).Name)
	}
	return res
}

func To_struct_ptr(obj interface{}) interface{} {
	return reflect.ValueOf(obj).Interface()
	/*
		vp := reflect.New(reflect.TypeOf(obj))
		vp.Elem().Set(reflect.ValueOf(obj))
		return vp.Interface()
	*/
}
