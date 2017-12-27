package template

import (
	"fmt"
	"path"
	"github.com/flosch/pongo2"
	"github.com/rfancn/goy2h/htmlelem"
)

var TEMPLATE_ROOT = "asset"
var gTemplateSet *pongo2.TemplateSet

func init(){
	memLoader := NewMemoryTemplateLoader(Asset)
	gTemplateSet = pongo2.NewSet("memory", memLoader)
}

func Render(el htmlelem.HtmlElementer) []byte{
	templateFile := el.GetTemplateFile()

	currentElem := el.GetCurrentElem()
	if currentElem != nil {
		return RenderElem(templateFile, currentElem)
	}
	return RenderElem(templateFile, el)
}

func RenderElem(templateFile string, el htmlelem.HtmlElementer) []byte {
	templatePath := path.Join(TEMPLATE_ROOT, templateFile)
	t := pongo2.Must(gTemplateSet.FromFile(templatePath))

	output, err := t.ExecuteBytes(pongo2.Context{"elem": el})
	if err != nil {
		fmt.Println(err)
	}

	return output
}


