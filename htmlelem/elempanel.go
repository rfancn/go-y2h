package htmlelem

const PANEL_DEFAULT_STYLE = "default"
const PANEL_DEFAULT_CLASS = "panel panel-default"

var PANEL_STYLE_CLASSES = map[string]string{
	"default":   "panel panel-default",
	"primary":   "panel panel-primary",
	"success":   "panel panel-success",
	"info":      "panel panel-info",
	"warning":   "panel panel-warning",
	"danger":    "panel panel-danger",
}

type panel struct {
	FieldsetElem
	Header string
	Body []HtmlElementer
	Footer []HtmlElementer
}

func (self *panel) Init(template string, elemName string, elemValue interface{}) {
	self.preInit(self, template, elemName, elemValue)
	self.processPanelStyle()
	self.processPanelHeader()
	self.processPanelBody()
	self.postInit()
}

func (self *panel) processPanelStyle() {
	style := bsMapGetOrDefault(self.ElemAttrMap, "panel-style", PANEL_DEFAULT_STYLE)
	class := sMapGetOrDefault(PANEL_STYLE_CLASSES, style, PANEL_DEFAULT_CLASS)

	delete(self.ElemAttrMap, "panel-style")
	self.insertCssClass(0, class)
}

func (self *panel) processPanelHeader() {
	self.Header = getElemStringAttr(self.ElemValue, "header")
}

func (self *panel) processPanelBody() {
	self.Body = getElemChildren(self.ElemValue, "body", self.Template)
}

func (self *panel) processPanelFooter() {
	self.Footer = getElemChildren(self.ElemValue, "footer", self.Template)
}



