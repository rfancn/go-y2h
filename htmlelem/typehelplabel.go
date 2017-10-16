package htmlelem

type HelpLabelElem struct{
	BaseElem
	HelpLabel map[string]string
}

func (self *HelpLabelElem) preInit(curElem HtmlElementer, template string, elemName string, elemValue interface{}) {
	self.BaseElem.preInit(curElem, template, elemName, elemValue)
	self.processHelpLabel()
}

func (el *HelpLabelElem) processHelpLabel() {
	// check if attribute map has key 'help-label' or not
	bsHelpLabel, ok := el.ElemAttrMap["help-label"]
	if !ok {
		return
	}

	mapHelpLabel := make(map[string]string)
	mapHelpLabel["innertext"] = string(bsHelpLabel)

	// get identifier and take it as "for" value
	// help-label.for used <label for="{{ help_label.for }}">
	bsId, ok := el.ElemAttrMap["id"]
	if ok {
		mapHelpLabel["for"] = string(bsId)
	}else {
		name, ok := el.ElemAttrMap["name"]
		if ok {
			mapHelpLabel["for"] = string(name)
		}
	}

	el.HelpLabel = mapHelpLabel
	delete(el.ElemAttrMap, "help-label")
}


