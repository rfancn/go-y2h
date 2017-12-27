package goy2h

import (
	"regexp"
	"github.com/rfancn/goy2h/htmlelem"
	"github.com/rfancn/goy2h/template"
)

const DEFAULT_TEMPLATE = "bootstrap3"

func (y2h *Y2H) GetHtml() string {
	htmlContent := make([]byte, 0)
	for _, elemValue := range y2h.yamlDocument.Html {
		el := htmlelem.NewElem(y2h.yamlDocument.Template, elemValue)
		// ignore those failed to initialize
		if el == nil {
			continue
		}

		//bsOutput := el.Render()
		bsOutput := template.Render(el)

		// try remove blank lines
		re := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z`)
		strOutput := re.ReplaceAllString(string(bsOutput), "")

		htmlContent = append(htmlContent, []byte(strOutput)...)
	}

	return string(htmlContent)
}
