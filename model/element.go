package model

type Element struct{
	name string
	Style
	value string
	parent element
}
// methods
func NewElement(name string, css Style, value string)*Element{
	return &Element{name, css, value, nil} 
}

func (this *Element) Render(){
	DrawElement(&this.Style, this.value, this.Parent())
}

func (this *Element) Focus(){
	MoveCursor(this.focus_x, this.focus_y)
}


func (this *Element) Name()string{
	return this.name
}

func (this *Element) GetStyle()*Style{
	return &this.Style
}

func (this *Element) IsEditable()bool{
	return false
}

func (this *Element) IsClickable()bool{
	return false
}

func (this *Element) Click(){}

func (this *Element)SetValue(v string)bool{
	this.value = v
	this.Render()
	return true
}

func (this *Element)GetValue()string{
	return this.value
}

func (this *Element)DelValue(){
	this.value = ""
}

func (this *Element) SetStatus(s int){}

func ( this *Element ) GetStatus()int{
	return 0
}

func (this *Element) Parent()element{
	return this.parent
}

func (this *Element) SetParent(e element){
	this.parent = e
}

func ( this *Element ) Select(name string)element{
	return nil
}

func (this *Element) GetState()bool{
	return false
}
func (this *Element) SetState(){}

func (this *Element) AddEvent(tipo string, f func()){}

