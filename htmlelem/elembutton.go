package htmlelem

const BUTTON_DEFAULT_STYLE = "default"
const BUTTON_DEFAULT_STYLE_CLASS = "btn btn-default"

var BTN_STYLE_CLASSES = map[string]string{
	"default":  "btn btn-default",
	"primary":  "btn btn-primary",
	"success":  "btn btn-success",
	"info":      "btn btn-info",
	"warning":  "btn btn-warning",
	"danger":   "btn btn-danger",
	"link":     "btn btn-link",
}

const BUTTON_DEFAULT_SIZE = "default"
const BUTTON_DEFAULT_SIZE_CLASS = ""

var BTN_SIZE_CLASSES = map[string]string{
	"default":  "",
	"small":    "btn-sm",
	"xsmall":   "btn-xs",
	"large":    "btn-lg",
}

type button struct {
	HelpLabelElem
	Text string
}

func (self *button) Init(template string, elemName string, elemValue interface{}) {
	self.preInit(self, template, elemName, elemValue)
	self.processButtonStyle()
	self.processButtonSize()
	self.processButtonText()
	self.postInit()
}

func (self *button) processButtonStyle() {
	style := bsMapGetOrDefault(self.ElemAttrMap, "button-style", BUTTON_DEFAULT_STYLE)
	class := sMapGetOrDefault(BTN_STYLE_CLASSES, style, BUTTON_DEFAULT_STYLE_CLASS)

	delete(self.ElemAttrMap, "button-style")
	self.insertCssClass(0, class)
}

func (self *button) processButtonSize() {
	size := bsMapGetOrDefault(self.ElemAttrMap, "size", BUTTON_DEFAULT_SIZE)
	class := sMapGetOrDefault(BTN_SIZE_CLASSES, size, BUTTON_DEFAULT_SIZE_CLASS)

	delete(self.ElemAttrMap, "size")
	//make sure size class after style class(whose length is 2)
	self.insertCssClass(2, class)
}

func (self *button) processButtonText() {
	self.Text = bsMapGetOrDefault(self.ElemAttrMap, "text", "")
	delete(self.ElemAttrMap, "text")
}
