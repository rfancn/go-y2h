package htmlelem

const RADIO_STYLE_DEFAULT = "default"
const RADIO_STYLE_INLINE = "inline"

const RADIO_CLASS_DEFAULT = "radio"
const RADIO_CLASS_INLINE = "radio-inline"

var RADIO_STYLE_CLASSES = map[string]string{
	RADIO_STYLE_DEFAULT: RADIO_CLASS_DEFAULT,
	RADIO_STYLE_INLINE:  RADIO_CLASS_INLINE,
}

type radio struct {
	ItemsElem
}

func (self *radio) Init(template string, elemName string, elemValue interface{}) {
	self.preInit(self, template, elemName, elemValue)
	// pay attention that processCheckboxStyle need to be invoked firstly,
	// otherwise, child item will inherit parent element's unused 'checkbox-style' attribute
	self.processRadioStyle()
	self.postInit()
}

func (self *radio) processRadioStyle() {
	var style string

	bsStyle, ok := self.ElemAttrMap["radio-style"]
	if !ok {
		style = RADIO_STYLE_DEFAULT
	}else {
		style = string(bsStyle)
		// if exist "checkbox-style", remove it to avoid it appears in AttributeString
		delete(self.ElemAttrMap, "radio-style")
	}

	class, ok := CHECKBOX_STYLE_CLASSES[style]
	if !ok {
		class = RADIO_CLASS_DEFAULT
	}

	self.insertCssClass(0, class)
}