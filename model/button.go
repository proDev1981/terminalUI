package model


// button
type button struct{
	Element
}

func Button(name string, css Style, value string, f func(Event))*button{
  return &button{Element{
    name:name, 
    Style:css, 
    value:value, 
    Listeners:Listeners{ 
      onClick:f, 
      onInside: func(e Event){}, 
      onOutside: func(e Event){}, 
    },
  }} 
}

func (this *button) IsClickable()bool{
	return true
}

