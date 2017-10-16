package htmlelem

const CHECKBOX_STYLE_DEFAULT = "default"
const CHECKBOX_STYLE_INLINE = "inline"

const CHECKBOX_CLASS_DEFAULT = "checkbox"
const CHECKBOX_CLASS_INLINE = "checkbox-inline"

var CHECKBOX_STYLE_CLASSES = map[string]string{
	CHECKBOX_STYLE_DEFAULT: CHECKBOX_CLASS_DEFAULT,
	CHECKBOX_STYLE_INLINE:  CHECKBOX_CLASS_INLINE,
}

type checkbox struct {
	ItemsElem
}

func (self *checkbox) Init(template string, elemName string, elemValue interface{}) {
	self.preInit(self, template, elemName, elemValue)
	self.processCheckboxStyle()
	self.postInit()
}

func (self *checkbox) processCheckboxStyle() {
	var style string

	bsStyle, ok := self.ElemAttrMap["checkbox-style"]
	if !ok {
		style = CHECKBOX_STYLE_DEFAULT
	}else {
		style = string(bsStyle)
		// if exist "checkbox-style", remove it to avoid it appears in AttributeString
		delete(self.ElemAttrMap, "checkbox-style")
	}

	class, ok := CHECKBOX_STYLE_CLASSES[style]
	if !ok {
		class = CHECKBOX_CLASS_DEFAULT
	}

	self.insertCssClass(0, class)
}