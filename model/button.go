package model


// button
type button struct{
	Element
	OnClick func(Event)
}

func Button(name string, css Style, value string, f func(Event))*button{
	return &button{Element{name, css, value, nil},f} 
}

func (this *button) IsClickable()bool{
	return true
}

func (this *button) Click(){
	this.OnClick(Event{"click", this})
	this.Focus()
}
