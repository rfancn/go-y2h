package htmlelem

const TABLE_DEFAULT_STYLE = "default"
const TABLE_DEFAULT_CLASS = "panel panel-default"

var TABLE_STYLE_CLASSES = map[string]string{
	"default":   "table",
	"striped":   "table table-striped",
	"bordered":   "table table-bordered",
	"hover":      "table table-hover",
	"condensed":   "table table-condensed",
}

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

func (self *table) processTableStyle() {
	style := bsMapGetOrDefault(self.ElemAttrMap, "table-style", PANEL_DEFAULT_STYLE)
	class := sMapGetOrDefault(PANEL_STYLE_CLASSES, style, PANEL_DEFAULT_CLASS)

	delete(self.ElemAttrMap, "table-style")
	self.insertCssClass(0, class)
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