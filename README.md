# go-y2h
y2h stands for: YAML to HTML, it aims to help translate YAML to HTML based on different templates.
It doesn't want to be a complete functional HTML generator, 
in most of time, it used as form component builder by only define some few lines.

### Introduction
Four item which can affect the translating of HTML can be defined in YAML document:
- template  
if no such item defines in YAML document, uses "bootstrap3" by default, now it only supports bootstrap3, it will support more templates in the future
  
- html  
html element definition, use HTML syntax
  
- javascript:  
There are 3 kinds of javascript can be defined here:
  - external:
    * src: specify the source url of the external javascript
    ```
    external: src="https://cdn.datatables.net/select/1.2.2/js/dataTables.select.min.js"
    ```
  - cdn:
    * locale:  
      specify locale of the CDN, it support "cn" and "en". when specify "cn", it will use bootcss CDN, if specify "en", it use cloudflare CDN  
      ```cdn: locale="cn"```
    * category:  
      specify the category of javascript library  
      ```cdn: name="select"```
    * ver:  
      specify the version of javascript library  
      ```cdn: ver="1.2.2"```
    * file:  
      specify the filename of javascript libray  
      ```cdn: file="js/dataTables.select.min.js"```
- css(under dev)
 
### Example
Example YAML document:
```yaml
html:
  - form: name="form1"
    fieldset:
    - input: help-label="input help label" type="text" value="input value" required
javascript:
  - inline: |
      function helloworld(){
        console.log("hello world!");
      }; 
```

If not specify template, it use "bootstrap3" as default template, and translate To:
#### HTML:
```html
<form name="form1" class="form-horizontal">
<div class="form-group">
    <label class="col-sm-3 control-label" for="">input help label</label>
    <div class="col-sm-9">
        <input class="form-control" type="text" required>
    </div>
</div>
</form>
```
#### Javascript:
Return Type: []map[string]string
```go
[
  map[string]string{"inline":"
    function helloworld(){
          console.log("hello world!");
    }; "
  },
  ...
]
```
