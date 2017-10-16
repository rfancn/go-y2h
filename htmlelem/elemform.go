package htmlelem

const LAYOUT_DEFAULT = "horizontal"
const LAYOUT_INLINE = "inline"
var FORM_LAYTOUT_CLASSES = map[string]string{
	LAYOUT_DEFAULT: "form-horizontal",
	LAYOUT_INLINE: "form-inline",
}

type form struct {
	FieldsetElem
	Layout string
}

func (self *form) Init(template string, elemName string, elemValue interface{}) {
	self.preInit(self, template, elemName, elemValue)
	self.processLayout()
	self.postInit()
}

func (el *form) processLayout() {
	var layout string

	// get defined layout
	bsLayout, ok := el.ElemAttrMap["layout"]
	if !ok {
		layout = LAYOUT_DEFAULT
	}else {
		layout = string(bsLayout)
	}

	// validate layout
	_, exist := FORM_LAYTOUT_CLASSES[layout]
	//only specified layout is the valid one and defined in FORM_LAYOUT_CLASSES keys
	if !exist {
		layout = LAYOUT_DEFAULT
	}

	delete(el.ElemAttrMap, "layout")
	el.insertCssClass(0, FORM_LAYTOUT_CLASSES[layout])
	el.Layout = layout
}



