package htmlelem

type table struct {
	BaseElem
	Thead []string
	Tbody [][]string
}

func (self *table) Init(template string, elemName string, elemValue interface{}) {
	self.preInit(self, template, elemName, elemValue)
	self.processTableThead()
	self.processTableTbody()
	self.postInit()
}

func (self *table) processTableThead() {
	//get []interface{} of thead
	self.Thead = getElemStringSliceAttr(self.ElemValue, "thead")
}

func (self *table) processTableTbody() {
	//get []interface{} of thead
	tbodySlice := getElemIFSliceAttr(self.ElemValue, "tbody")
	if tbodySlice == nil{
		return
	}

	self.Tbody = make([][]string, len(tbodySlice))
	for i, child := range tbodySlice {
		childSlice, ok := child.([]interface{})
		if !ok {
			continue
		}

		//convert []interface{} to []string
		self.Tbody[i] = IFSlice2StrSlice(childSlice)
	}
}