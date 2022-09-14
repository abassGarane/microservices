package data

import "testing"


func TestCheckValidation(t *testing.T){
  p := &Product{}
  err := p.Validate();

  if err != nil {
    t.Fatal(err)
  }
}
