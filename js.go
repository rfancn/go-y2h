package go_y2h

import (
	"fmt"
	"bytes"
	"log"
)

const CDN_LOCALE_CN = "cn"
const CDN_LOCALE_EN = "en"
const CDN_LOCALE_DEFAULT = CDN_LOCALE_CN

//<base url>/category/ver/file
const CDN_BOOTCSS_URL_TEMPLATE = "https://cdn.bootcss.com/%s/%s/%s"
const CDN_CLOUDFLARE_URL_TEMPLATE = "https://cdnjs.cloudflare.com/ajax/libs/%s/%s/%s"
const CDN_URL_TEMPLATE_DEFAULT = CDN_BOOTCSS_URL_TEMPLATE

const EXTERNAL_JS = "<script src=\"%s\"></script>"

var CDN_URL_TEMPLATES = map[string]string{
	CDN_LOCALE_CN: CDN_BOOTCSS_URL_TEMPLATE,
	CDN_LOCALE_EN: CDN_CLOUDFLARE_URL_TEMPLATE,
}

func (y2h *FileY2H) GetJavascript() map[string]string {
	var inlineJsBuffer bytes.Buffer
	var externalJsBuffer bytes.Buffer

	for _, jsValue := range y2h.yamlDocument.Javascript {
		//jsValue must be a map
		jsMap, ok := jsValue.(map[interface{}]interface{})
		if !ok {
			continue
		}

		for k,v := range jsMap {
			jsType := fmt.Sprint(k)
			switch jsType {
				case "cdn":
					externalJsBuffer.WriteString(getCdnJS(v))
					externalJsBuffer.WriteString("\n")
				case "external":
					externalJsBuffer.WriteString(getExternalJS(v))
					externalJsBuffer.WriteString("\n")
				case "inline":
					inlineJsBuffer.WriteString(getInlineJS(v))
					inlineJsBuffer.WriteString("\n")
			}
		}
	}

	jsSlice := make(map[string]string)
	jsSlice["inline"] = inlineJsBuffer.String()
	jsSlice["external"] = externalJsBuffer.String()

	return jsSlice
}

func getInlineJS(jsValue interface{}) string {
	// if it is inline, then jsValue should be a string
	jsContent, ok := IF2String(jsValue)
	if !ok {
		log.Println("getInlineJs():Type Assertion to string failed")
		return ""
	}

	return jsContent
}

func getCdnJS(jsValue interface{}) string {
	// if it is inline, then jsValue should be a string
	cdnString, ok := IF2String(jsValue)
	if !ok {
		log.Println("getCdnJS():Type Assertion to string failed")
		return ""
	}

	//convert cdn string to cdn map
	cdnMap := convertKVStringToMap(cdnString)

	//validate if cndMap contains required keys
	var requireKeys = []string{"category", "ver", "file"}
	for _, k := range requireKeys {
		if _, exist := cdnMap[k]; !exist {
			log.Printf("getCdnJS():Required key:%s doesn't exist\n", k)
			return ""
		}
	}

	cdnLocal, exist := cdnMap["locale"]
	if !exist {
		cdnLocal = CDN_LOCALE_DEFAULT
	}

	cdnUrlTemplate, exist := CDN_URL_TEMPLATES[cdnLocal]
	if !exist {
		cdnUrlTemplate = CDN_URL_TEMPLATE_DEFAULT
	}

	cdnURL := fmt.Sprintf(cdnUrlTemplate, cdnMap["category"], cdnMap["ver"], cdnMap["file"])

	return buildExternalJavascript(cdnURL)
}

func getExternalJS(jsValue interface{}) string {
	// if it is inline, then jsValue should be a string
	//e,g: src="https://cdn.datatables.net/select/1.2.2/js/dataTables.select.min.js"
	jsContent, ok := IF2String(jsValue)
	if !ok {
		log.Println("getExternalJS():Type Assertion to string failed")
		return ""
	}

	//convert cdn string to cdn map
	jsMap := convertKVStringToMap(jsContent)
	//validate if jsMap contains required "src" key
	jsSrc, exist := jsMap["src"]
	if !exist {
		log.Println("getExternalJS():Required key \"src\" doesn't exist")
		return ""
	}

	return buildExternalJavascript(jsSrc)
}

func IF2String(value interface{}) (string, bool) {
	strValue, ok := value.(string)
	if !ok {
		return "", false
	}
	return strValue, true
}

func buildExternalJavascript(src string) string {
	return fmt.Sprintf(EXTERNAL_JS, src)
}
