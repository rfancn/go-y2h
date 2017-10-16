package htmlelem

import (
	"reflect"
	"fmt"
	"log"
)

//central registry of html element types
var elementTypeRegistry = make(map[string]reflect.Type)
// save all registered html element types set
var elementTypeSet = make(map[string]uint8)

type HtmlElementer interface {
	Init(template string, elemName string, elemValue interface{})
	Render() []byte
}

func init() {
	// register supported element type with specific struct
	// This way will reduce size of allocations when building the map
	// when compared with registerElemType(&form{})
	registerElemType(&form{})
	registerElemType(&input{})
	registerElemType(&button{})
	registerElemType(&item{})
	registerElemType(&checkbox{})
	registerElemType(&radio{})
	registerElemType(&table{})
	registerElemType(&panel{})
	registerElemType(&modal{})

	// record all supported element types
	// so we can use it in getElemName function
	for k, _ := range elementTypeRegistry {
		elementTypeSet[k] = 1
	}
}

func registerElemType(elem interface{}) {
	t := reflect.TypeOf(elem).Elem()
	elementTypeRegistry[t.Name()] = t
}

//validateElemName check if a string contains in ElementTypeSet
// if true, then elemName is valid
func validateElemName(v reflect.Value) (elemName string, ok bool) {
	// try convert the interface{} to string
	guessName, ok := v.Interface().(string)
	if !ok {
		return "", false
	}

	// if found the key in current supported element type set
	if _, ok := elementTypeSet[guessName]; !ok {
		return "", false
	}

	return guessName, true
}

//guessElemName will check following:
// 1. extract elemName from the elemValue
// 2. extract elemType from the elemValue
func guessElemName(elemValue interface{}) (elemName string, ok bool) {
	v := reflect.ValueOf(elemValue)
	switch kind := v.Kind(); kind {
		case reflect.String:
			return validateElemName(v)
		case reflect.Map:
			for _, key := range v.MapKeys() {
				if elemName, ok := validateElemName(key); ok{
					return elemName, ok
				}
			}
		default:
			log.Println("Invalid element: it should be a string or map")
	}

	return "", false
}

//NewElem create HtmlElement instance
func NewElem(template string, elemValue interface{})  HtmlElementer {
	elemName, ok := guessElemName(elemValue)
	if !ok {
		fmt.Println("Failed to guess element name")
		return nil
	}

	// vPointOfElem is reflect.Value of a pointer to the specific element
	vPointOfElem := reflect.New(elementTypeRegistry[elemName])

	// refelct.Value.Interface() convert reflect.Value to interface{}
	// then type assertion will work because interface can be either of point to struct or struct value
	// After converting, pElem is a point to specific element struct
	// e,g: var elem *htmlelem.form = &form{}
	pElem, ok:= vPointOfElem.Interface().(HtmlElementer)
	if !ok {
		log.Println("Failed to convert to HtmlElementer interface!")
		return nil
	}

	pElem.Init(template, elemName, elemValue)

	return pElem
}




/**
//NewByReflect() will create a Html Element instance based on template, element name, and element value
//the corresponding element struct get by elemName as key from center type registry
//it does following:
//1. guess element name from elemant value
//2. try to get this element's attribute map
//3. set elemName, elemValue, emeAttrMap to specific element instance
func NewByReflect(template string, elemValue interface{})  HtmlElementer {
	elemName, ok := guessElemName(elemValue)
	if !ok {
		fmt.Println("Failed to guess element name")
		return nil
	}

	// vPointOfElem is reflect.Value of a pointer to the specific element
	vPointOfElem := reflect.New(elementTypeRegistry[elemName])

	//The reflect way to setup initialize value for instance
	//no Init() need to be called for each instance
	if  !setElemField(vPointOfElem, "Template",  template) ||
		!setElemField(vPointOfElem, "ElemName",  elemName) ||
		!setElemField(vPointOfElem, "ElemValue", elemValue) ||
		!setElemField(vPointOfElem, "ElemAttrMap", buildElemAttrMap(elemName, elemValue)) {
		log.Println("Failed to initialize element")
		return nil
	}

	// refelct.Value.Interface() convert reflect.Value to interface{}
	// then type assertion will work because interface can be either of point to struct or struct value
	// After converting, pElem is a point to specific element struct
	// e,g: var elem *htmlelem.form = &form{}
	pElem, ok:= vPointOfElem.Interface().(HtmlElementer)
	if !ok {
		log.Println("Failed to convert to HtmlElementer interface!")
		return nil
	}

	return pElem
}

func setElemField(pElem reflect.Value, fieldName string, fieldValue interface{}) bool {
	pData := pElem.Elem()

	if pData.Kind() == reflect.Struct {
		// exported field
		field := pData.FieldByName(fieldName)
		if field.IsValid() {
			// A Value can be changed only if it is addressable
			// and was not obtained by the use of unexported struct fields.
			if field.CanSet() {
				// set value based on field.Kind()
				switch kind := field.Kind(); kind {
				// used for set Template and ElemType
				case reflect.String:
					v := fmt.Sprintf("%v", fieldValue)
					field.SetString(v)
					return true
				case reflect.Int:
					if v, ok := fieldValue.(int64); ok {
						field.SetInt(v)
						return true
					}
					// used for set ElemValue
				case reflect.Interface: fallthrough
				case reflect.Map:
					field.Set(reflect.ValueOf(fieldValue))
					return true
				default:
					log.Printf("Unsupported set field:[%s]", kind)
					return false
				}
			}
		}
	}

	return false
}
**/

