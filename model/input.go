package model


// input
type input struct{
	Element
}

func Input(name string, css Style, value, holder string)*input{
  res := &input{Element{name:name, Style:css, value:value, Listeners:DEFAULT_LISTENERS}}
  res.Style.Playholder = holder
  return res
}

func ( i *input ) IsEditable()bool{
	return true
}

func (i *input) SetValue(v string)bool{
	if len(i.Element.value) < i.Width-1 {
		i.Element.value += v
	}else{
		return false
	}
	return true
}

func (i *input) Focus(){
  i.Element.Focus()
	MoveCursor(i.focus_x + len(i.value), i.focus_y)
}
