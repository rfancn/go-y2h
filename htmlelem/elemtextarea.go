package htmlelem

type textarea struct {
	HelpLabelElem
}

func (self *textarea) Init(template string, elemName string, elemValue interface{}) {
	self.preInit(self, template, elemName, elemValue)
	self.postInit()
}