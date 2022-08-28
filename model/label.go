package model


// Label
type label struct{
	Element
}

func Label(name string, css Style, value string)*label{
  return &label{Element{name:name, Style:css, value:value, Listeners:DEFAULT_LISTENERS}} 
}

