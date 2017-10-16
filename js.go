package go_y2h

import (
	"fmt"

)

const CDN_LOCALE_CN = "cn"
const CDN_LOCALE_EN = "en"
const CDN_LOCALE_DEFAULT = CDN_LOCALE_CN

//<base url>/category/ver/file
const CDN_BOOTCSS_URL_TEMPLATE = "https://cdn.bootcss.com/%s/%s/%s"
const CDN_CLOUDFLARE_URL_TEMPLATE = "https://cdnjs.cloudflare.com/ajax/libs/%s/%s/%s"
const CDN_URL_TEMPLATE_DEFAULT = CDN_BOOTCSS_URL_TEMPLATE

const EXTERNAL_JS = "<script src=%s></script>"

var CDN_URL_TEMPLATES = map[string]string{
	CDN_LOCALE_CN: CDN_BOOTCSS_URL_TEMPLATE,
	CDN_LOCALE_EN: CDN_CLOUDFLARE_URL_TEMPLATE,
}

func (y2h *FileY2H) GetJavascript() []map[string]string {
	jsSlice := make([]map[string]string, 0)

	for _, jsValue := range y2h.yamlDocument.Javascript {
		//jsValue must be a map
		jsMap, ok := jsValue.(map[interface{}]interface{})
		if !ok {
			continue
		}

		var js map[string]string
		for k,v := range jsMap {
			jsType := fmt.Sprint(k)
			switch jsType {
				case "cdn":
					js = getCdnJS(v)
				case "inline":
					js = getInlineJS(v)
				case "external":
					js = getExternalJS(v)
			}
			if js != nil {
				jsSlice = append(jsSlice, js)
			}
		}
	}

	return jsSlice
}

func getInlineJS(jsValue interface{}) map[string]string {
	// if it is inline, then jsValue should be a string
	jsContent, ok := IF2String(jsValue)
	if !ok {
		return nil
	}

	js := make(map[string]string)
	js["inline"] = jsContent

	return js
}

func getCdnJS(jsValue interface{}) map[string]string {
	// if it is inline, then jsValue should be a string
	cdnString, ok := IF2String(jsValue)
	if !ok {
		return nil
	}

	//convert cdn string to cdn map
	cdnMap := convertKVStringToMap(cdnString)

	//validate if cndMap contains required keys
	var requireKeys = []string{"category", "ver", "file"}
	for _, k := range requireKeys {
		if _, exist := cdnMap[k]; !exist {
			return nil
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

func getExternalJS(jsValue interface{}) map[string]string {
	// if it is inline, then jsValue should be a string
	jsSrc, ok := IF2String(jsValue)
	if !ok {
		return nil
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

func buildExternalJavascript(src string) map[string]string {
	js := make(map[string]string)
	js["external"] = fmt.Sprintf(EXTERNAL_JS, src)

	return js
}
