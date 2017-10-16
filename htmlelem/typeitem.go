package htmlelem

type item struct {
	AttrMap map[string][]byte
	Parent HtmlElementer
	Attribute string
	Label string
}

/**
func getItemAttrMap(parent *ItemsElem, itemValue interface{}) map[string][]byte {
	itemAttrMap := make(map[string][]byte)

	//deep copy parent attributes to child item attributes
	for k,v := range parent.ElemAttrMap {
		//ignore parent items's class attribute
		if k == "class" {
			continue
		}
		itemAttrMap[k] = v
	}
	//get original item attributes
	origItemAttrMap := buildElemAttrMap("item", itemValue)
	// currently itemAttrMap is a copy of parent.ElemAttrMap
	// combine item attribute with parent element's attribute
	for k,v := range origItemAttrMap {
		//ignore specific label attribute
		if k == "label" {
			continue
		}
		//append item attribute that not appears in parent element attribute
		_, exist := itemAttrMap[k]
		if !exist {
			itemAttrMap[k] = v
		}
	}

	return itemAttrMap
}
**/

func getItemLabel(itemAttrMap map[string][]byte) string {
	sbLabel, exist := itemAttrMap["item-label"]
	if !exist {
		return ""
	}

	//remove specific attribute "item-label", so it will not appears in attribute string
	delete(itemAttrMap, "item-label")
	return string(sbLabel)
}

//NewChildElem create HtmlElement instance that has child reference
func NewItem(parent *ItemsElem, itemValue interface{})  *item {
	itemAttrMap := buildElemAttrMap("item", itemValue)

	pItem := &item{}
	pItem.Parent = parent
	pItem.Label = getItemLabel(itemAttrMap)
	pItem.Attribute = getElemAttrString(itemAttrMap)

	return pItem
}

