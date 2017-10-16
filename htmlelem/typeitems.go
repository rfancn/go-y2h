package htmlelem

//HtmlElement with HelpLabel and Items
type ItemsElem struct {
	HelpLabelElem
	Items []*item
}

func (self *ItemsElem) preInit(curElem HtmlElementer, template string, elemName string, elemValue interface{}) {
	self.HelpLabelElem.preInit(curElem, template, elemName, elemValue)
	self.processItems()
}

//Parse child items for radio and checkbox element
func (self *ItemsElem) processItems() {
	self.Items = make([]*item, 0)

	// elemValue must be a map
	elemValueMap, ok := self.ElemValue.(map[interface{}]interface{})
	if !ok {
		//in case, no items nor attribute defines for radio element, ElemValue is a string
        //this is a rare case but still valid, check it here
        //e,g: - form:
        //        - radio
		return
	}

	items, ok := elemValueMap["items"]
	// if no "items" defined
	if !ok {
		return
	}

	// items must be a slice
	itemSlice, ok := items.([]interface{})
	if !ok {
		return
	}

	// create child item instance
	for _, item := range itemSlice {
		newItem := NewItem(self, item)

		// ignore those failed to initialize
		if newItem == nil {
			continue
		}

		self.Items = append(self.Items, newItem)
	}
}

