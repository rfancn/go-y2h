package y2h

import (
    "testing"
)

var y = New()

func TestEmptyForm(t *testing.T) {

   if ok := y.Read("examples/form.yaml"); !ok{
     t.Error("Failed to read test.yaml")
   }

   htmlContent := y.GetHtml()
   if len(htmlContent) == 0 {
       t.Error("Failed to render form html element")
   }
}

func TestButtons(t *testing.T) {
    y := New()
    if ok := y.Read("examples/form_button.yaml"); !ok{
        t.Error("Failed to read test.yaml")
    }

    htmlContent := y.GetHtml()
    if len(htmlContent) == 0 {
        t.Error("Failed to render form button html element")
    }
}
