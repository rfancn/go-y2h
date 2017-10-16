package htmlelem

type modal struct {
	BaseElem
	Header string
	Body []HtmlElementer
	Footer []HtmlElementer
}

func (self *modal) Init(template string, elemName string, elemValue interface{}) {
	self.preInit(self, template, elemName, elemValue)
	self.processModalHeader()
	self.processModalBody()
	self.processModalFooter()
	self.postInit()
}

func (self *modal) processModalHeader() {
	self.Header = getElemStringAttr(self.ElemValue, "header")
}

func (self *modal) processModalBody() {
	self.Body = getElemChildren(self.ElemValue, "body", self.Template)
}

func (self *modal) processModalFooter() {
	self.Footer = getElemChildren(self.ElemValue, "footer", self.Template)
}

