package htmlelem

type FieldsetElem struct {
	BaseElem
	Fieldset []HtmlElementer
}

func (self *FieldsetElem) preInit(curElem HtmlElementer, template string, elemName string, elemValue interface{}) {
	self.BaseElem.preInit(curElem, template, elemName, elemValue)
	self.processFieldset()
}

func (el *FieldsetElem) processFieldset() {
	el.Fieldset = make([]HtmlElementer, 0)

	fieldsetSlice := getElemIFSliceAttr(el.ElemValue, "fieldset")

	/**
	// elemValue must be a map
	elemValueMap, ok := el.ElemValue.(map[interface{}]interface{})
	if !ok {
		return
	}

	//if fieldset defined
	fieldset, ok := elemValueMap["fieldset"]
	if !ok {
		return
	}

	// fieldset must be a slice
	fieldsetSlice, ok := fieldset.([]interface{})
	if !ok {
		return
	}
	**/

	for _, childElemValue := range fieldsetSlice {
		childElem := NewElem(el.Template, childElemValue)
		// ignore those failed to initialize
		if childElem == nil {
			continue
		}

		el.Fieldset = append(el.Fieldset, childElem)
	}
}
