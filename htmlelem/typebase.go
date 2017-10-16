package htmlelem

import (
	"fmt"
	"strings"
	"path"
	"github.com/anmitsu/go-shlex"
	"github.com/flosch/pongo2"
)

type BaseElem struct{
	Template  string
	ElemName  string
	ElemValue interface{}
	ElemAttrMap map[string][]byte

	Attribute string
	// points to the real elem instance
	Current HtmlElementer
}

func (el *BaseElem) preInit(curElem HtmlElementer, template string, elemName string, elemValue interface{}) {
	//invoke base Init to initialize all variables
	el.Init(template, elemName, elemValue)
	// set current element to who invoke super()
	el.Current = curElem

	/**
	// handle specific attribute
	specAttrSlice := make([]string, 0)
	for specAttrName, specAttrProcessfunc := range el.SpecificAttrMap {
		// call specific attr processing func
		specAttrProcessfunc()
		specAttrSlice = append(specAttrSlice, specAttrName)
	}
	**/

}

func (el *BaseElem) postInit() {
	el.Attribute = getElemAttrString(el.ElemAttrMap)
}

//simulate super.__init__() syntax in Python
func (el *BaseElem) Init(template string, elemName string, elemValue interface{}) {

	//set TemplateFile
	el.Template = template
	el.ElemName = elemName
	el.ElemValue = elemValue

	el.ElemAttrMap = buildElemAttrMap(elemName, elemValue)

	el.Attribute = ""
	el.Current = nil
}

func (el *BaseElem) Render() []byte{
	//get template file
	file := fmt.Sprintf("%s.html", el.ElemName)
	templateFile := path.Join("goy2h", "templates", el.Template, file)

	if el.Current != nil {
		return RenderElem(templateFile, el.Current)
	}
	return RenderElem(templateFile, el)
}

//parseElemAttrStr parse element attribute string to map
// e,g: input: name="abc" value="123" required
// map[name:abc value:123 required:nil]
func parseAttrStr(attrStr string) map[string][]byte {
	elemAttrMap := make(map[string][]byte)

	// slice to save single attribute, e,g: 'required ' in `input: name="abc" value="123" required`
	var singleAttrSlice []string
	// map to save pairs attribute, e,g: name="abc"
	pairsAttrMap := make(map[string]string)

	sList, _ := shlex.Split(attrStr, true)
	for _, word := range sList {
		equalsSignLoc := strings.Index(word, "=")

		// handle single attribute
		if equalsSignLoc == -1 {
			singleAttrSlice = append(singleAttrSlice, word)
			continue
		}

		// handle pairs attribute
		k := word[:equalsSignLoc]
		v := word[equalsSignLoc+1:]
		pairsAttrMap[k] = v
	}

	if len(singleAttrSlice) == 0 && len(pairsAttrMap) == 0 {
		return elemAttrMap
	}


	// add double attr to elemAttrMap
	for k,v := range pairsAttrMap {
		// remove empty space in value
		s := string(v)
		s = strings.Trim(s, " ")
		elemAttrMap[k] = []byte(s)
	}

	// add single attr into elem attribute map
	for _, singleAttr := range singleAttrSlice {
		elemAttrMap[singleAttr] = nil
	}

	return elemAttrMap
}

func RenderElem(templateFile string, el HtmlElementer) []byte {
	t := pongo2.Must(pongo2.FromFile(templateFile))

	output, err := t.ExecuteBytes(pongo2.Context{"elem": el})
	if err != nil {
		fmt.Println(err)
	}

	return output
}

//getElemAttrString convert attribute Map to attribute string
func getElemAttrString(elemAttrMap map[string][]byte) string {
	kvSlice := make([]string, 0)
	singleSlice := make([]string, 0)
	for k,v := range elemAttrMap {
		if v != nil {
			attr := fmt.Sprintf(`%s="%s"`, k, v)
			kvSlice = append(kvSlice, attr)
		}else{
			singleSlice = append(singleSlice, k)
		}
	}

	attrSlice := append(kvSlice, singleSlice...)
	return strings.Join(attrSlice, " ")
}

func (el *BaseElem) insertCssClass(index int, newClassStr string) {
	newClassFields := strings.Fields(newClassStr)
	if len(newClassFields) == 0 {
		return
	}

	origClassByteSlice, ok := el.ElemAttrMap["class"]
	// if no original class has been set, then set to be new one
	if !ok {
		el.ElemAttrMap["class"] = []byte(newClassStr)
		return
	}

	origClassFields := strings.Fields(string(origClassByteSlice))

	insertFields := make([]string, 0)
	for _, newClassField := range newClassFields {
		if ! stringInSlice(newClassField, origClassFields){
			insertFields = append(insertFields, newClassField)
		}
	}

	finalClassFields := make([]string, 0)
	if index == 0 {
		finalClassFields = append(insertFields, origClassFields...)
	}else if index >= len(origClassFields) {
		finalClassFields = append(origClassFields, insertFields...)
	}else {
		finalClassFields = append(origClassFields[:index], append(insertFields, origClassFields[index:]...)...)
	}

	el.ElemAttrMap["class"] = []byte(strings.Join(finalClassFields, " "))
}
