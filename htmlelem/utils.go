package htmlelem

import (
	"reflect"
	"strings"
	"sort"
	"fmt"
)

//bsMapGetOrDefault try to get value by key in map[string][]byte
//if no such key, return defaultValue, otherwise return string(value)
func bsMapGetOrDefault(dict map[string][]byte, key string, defaultValue string) string {
	value, ok := dict[key]
	if !ok {
		return defaultValue
	}

	return string(value)
}

func sMapGetOrDefault(dict map[string]string, key string, defaultValue string) string {
	value, ok := dict[key]
	if !ok {
		return defaultValue
	}

	return value
}

//buildElemAttrMap will parse string after the elemName and convert it to map[string]string
//e,g: form: name="form1" layout="inline" => map{name:form1 layout:inline}
//There are serveral cases when define HtmlElement in YAML
// 1. - form
// 2. - form:
// 3. - form: name="form1"
// for case 1, the elemValue is a string "form"
// for case 2, the elemValue is a map with emtpy value, map{form:}
// for case 3, the elemValue is a map with attribute string as value, map[form:'attribute string']
func buildElemAttrMap(elemName string, elemValue interface{}) map[string][]byte {
	elemAttrMap := make(map[string][]byte)

	v := reflect.ValueOf(elemValue)
	kind := v.Kind()

	// case 1: no attribute string for elemName, even empty string
	if kind == reflect.String {
		return elemAttrMap
	}else if kind == reflect.Map {
		attrStr := getElemStringAttr(elemValue, elemName)
		// in case, there is empty attribute string after elemName
		attrStr = strings.Trim(attrStr, " ")
		if len(attrStr) == 0 {
			return elemAttrMap
		}
		return parseAttrStr(attrStr)
	}

	return elemAttrMap
}


func stringInSlice(a string, strlist []string) bool {
	sort.Strings(strlist)
	listLength := len(strlist)
	i := sort.Search(listLength, func(i int) bool { return strlist[i] >= a })
	if i < listLength && strlist[i] == a {
		return true
	}
	return false
}

//getElemAttrValue returns the interface{} representation of attribute Value
func getElemAttrValue(elemValue interface{}, attrName string) interface{} {
	value := reflect.ValueOf(elemValue)
	if value.Kind() != reflect.Map {
		return nil
	}

	// find the map value by elemName key and convert it from reflect.Value to interface{}
	attrValue := value.MapIndex(reflect.ValueOf(attrName))
	if !attrValue.IsValid() {
		return nil
	}

	return attrValue.Interface()
}

//getElemStringAttr returns element's attribute which represent as string
func getElemStringAttr(elemValue interface{}, attrName string) string {
	attrValue := getElemAttrValue(elemValue, attrName)
	if attrValue == nil {
		return ""
	}

	attrString, ok := attrValue.(string)
	if !ok {
		return ""
	}

	return attrString
}

//getElemIFSliceAttr returns element's attribute which represent as []interface{}
func getElemIFSliceAttr(elemValue interface{}, attrName string) []interface{} {
	attrValue := getElemAttrValue(elemValue, attrName)

	attrSlice, ok := attrValue.([]interface{})
	if !ok {
		return nil
	}

	return attrSlice
}

func getElemChildren(elemValue interface{}, childKey string, template string) []HtmlElementer {
	childElems := make([]HtmlElementer, 0)

	childSlice := getElemIFSliceAttr(elemValue, childKey)

	for _, childElemValue := range childSlice {
		childElem := NewElem(template, childElemValue)
		// ignore those failed to initialize
		if childElem == nil {
			continue
		}

		childElems = append(childElems, childElem)
	}

	return childElems
}

//getElemIFSliceAttr returns element's attribute which represent as []interface{}
func getElemStringSliceAttr(elemValue interface{}, attrName string) []string {
	ifsAttrValue := getElemIFSliceAttr(elemValue, attrName)
	return IFSlice2StrSlice(ifsAttrValue)
}

//IFSlice2StrSlice convert []interface{} to []string
func IFSlice2StrSlice(is []interface{}) []string {
	ss := make([]string, len(is))
	for i, v := range is {
		ss[i] = fmt.Sprint(v)
	}
	return ss
}
