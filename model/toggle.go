package model


type toggle struct{
	Element
	label string
	state bool
	OnClick func(e Event)
}

func Toggle(name string, s Style, label string, f func(Event))*toggle{
	return &toggle{Element{name, s, "☐", nil},label, false, f}
}

func (this *toggle) Render(){
	DrawElement(&this.Style, this.value  +  this.label, this.Parent())
}

func ( this *toggle ) IsClickable()bool{
	return true
}

func ( this *toggle ) GetState()bool{
	return this.state
}
func (this *toggle) SetState(){
	this.state = !this.state
	if this.state{ 
		this.SetValue("☑" + this.label ) 
	}else{
		this.SetValue("☐" + this.label )
	}
}

func (this *toggle) Click(){
	this.SetState()
	this.OnClick(Event{"click", this})
}
