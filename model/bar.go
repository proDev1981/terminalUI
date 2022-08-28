package model

import "strings"

type statusBar struct{
	Element
	status int
	progres string
}

func Bar(name string, s Style)*statusBar{
  return &statusBar{Element:Element{ name:name, Style:s, Listeners:DEFAULT_LISTENERS}}
}

func (this *statusBar) Render(){
	this.progres = strings.Repeat(" ", (this.Style.Width*this.status/100))
	DrawElement(&this.Style, this.progres, this.Parent())
}

func (this *statusBar) GetStatus()int{
	return this.status
}

func (this *statusBar) SetStatus(s int){
	if len(this.progres) < this.Style.Width{
		this.status = s
		this.Render()
	}
}
