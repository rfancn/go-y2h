package htmlelem

type input struct {
	HelpLabelElem
}

func (self *input) Init(template string, elemName string, elemValue interface{}) {
	self.preInit(self, template, elemName, elemValue)
	self.postInit()
}