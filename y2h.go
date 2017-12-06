//Package goy2h try to convert YAML to HTML
package y2h

import (
	"io/ioutil"
	"fmt"
	"strings"
	"gopkg.in/yaml.v2"
	"github.com/anmitsu/go-shlex"
)

type YAMLDocument struct {
	Template string             `yaml:"template"`
	Html       []interface{}   `yaml:"html"`
	Javascript []interface{}   `yaml:"javascript"`
}

type FileY2H struct{
	yamlDocument *YAMLDocument
}

func New() *FileY2H {
	return &FileY2H{}
}

func (y2h *FileY2H) Read(yamlFilename string) bool {
	yamlDocument, err := parseYaml(yamlFilename)
	if err != nil {
		fmt.Println(err)
		return false
	}

	y2h.yamlDocument = yamlDocument
	return true
}

//parseYaml and returns YamlDocument instance
func parseYaml(yamlFilename string) (*YAMLDocument, error) {
	// read yaml file and unmarshal to YAMLDocument struct
	content, err := ioutil.ReadFile(yamlFilename)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to read YAML file: %s", yamlFilename)
		return nil, NewY2HError(errMsg, err)
	}

	//if len(content) <=0 {
	//	return NewY2HError("Empty YAML content")
	//}
	yamlDocument := &YAMLDocument{}
	err = yaml.Unmarshal(content, yamlDocument)
	if err != nil {
		return nil, NewY2HError("Failed to unmarshal YAML", err)
	}

	//handle tempalate
	if yamlDocument.Template == "" {
		yamlDocument.Template = DEFAULT_TEMPLATE
	}

	return yamlDocument, nil
}

//convertKVStringToMap parse key/value string to map
// e,g: name="abc" value="123" required
// map[name:abc value:123]
func convertKVStringToMap(kvString string) map[string]string {
	dict := make(map[string]string)

	kvSlice, _ := shlex.Split(kvString, true)
	for _, word := range kvSlice {
		equalsSignLoc := strings.Index(word, "=")

		// if failed to locate "=" symbol in word,
		// then this is a invalid key/value word
		if equalsSignLoc == -1 {
			continue
		}

		// handle pairs attribute
		k := string(word[:equalsSignLoc])
		v := string(word[equalsSignLoc+1:])
		// remove empty space in value
		v = strings.Trim(v, " ")
		dict[k] = v
	}

	return dict
}

