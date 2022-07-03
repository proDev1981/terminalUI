package model


// input
type input struct{
	Element
}

func Input(name string, css Style, value string)*input{
	return &input{Element{name, css, value, nil}}
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
	MoveCursor(i.focus_x + len(i.value), i.focus_y)
}
