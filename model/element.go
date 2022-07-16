package model

import "strings"
import "strconv"

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

func (this *Element) IsBox()bool{
	return false
}

func (this *Element) IsFocusable()bool{
	return 	this.IsEditable() || 
			this.IsClickable() || 
			this.IsBox()  
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

func (this *Element) GetOrder()int{
	return this.Style.order
}

func (this *Element) SetOrder(order int){
	this.Style.order = order
}

func (this *Element) GetMarginLeft()(res int){
	strMargin := strings.Split(this.Style.Margin, ",")
	switch len(strMargin){
	case 0:
		res = 0
	case 1:
		res, _ = strconv.Atoi(strMargin[0])
	case 2:
		res, _ = strconv.Atoi(strMargin[0])
	case 3:
		res, _ = strconv.Atoi(strMargin[0])
	case 4:
		res, _ = strconv.Atoi(strMargin[0])
	default:
		res, _ = strconv.Atoi(strMargin[0])
	}
	return
}
func (this *Element) GetMarginRight()(res int){
	strMargin := strings.Split(this.Style.Margin, ",")
	switch len(strMargin){
	case 0:
		res = 0
	case 1:
		res, _ = strconv.Atoi(strMargin[0])
	case 2:
		res, _ = strconv.Atoi(strMargin[0])
	case 3:
		res, _ = strconv.Atoi(strMargin[1])
	case 4:
		res, _ = strconv.Atoi(strMargin[1])
	default:
		res, _ = strconv.Atoi(strMargin[0])
	}
	return
}
func (this *Element) GetMarginBotton()(res int){
	strMargin := strings.Split(this.Style.Margin, ",")
	switch len(strMargin){
	case 0:
		res = 0
	case 1:
		res, _ = strconv.Atoi(strMargin[0])
	case 2:
		res, _ = strconv.Atoi(strMargin[1])
	case 3:
		res, _ = strconv.Atoi(strMargin[2])
	case 4:
		res, _ = strconv.Atoi(strMargin[3])
	default:
		res, _ = strconv.Atoi(strMargin[0])
	}
	return
}
func (this *Element) GetMarginTop()(res int){
	strMargin := strings.Split(this.Style.Margin, ",")
	switch len(strMargin){
	case 0:
		res = 0
	case 1:
		res, _ = strconv.Atoi(strMargin[0])
	case 2:
		res, _ = strconv.Atoi(strMargin[1])
	case 3:
		res, _ = strconv.Atoi(strMargin[2])
	case 4:
		res, _ = strconv.Atoi(strMargin[2])
	default:
		res, _ = strconv.Atoi(strMargin[0])
	}
	return
}
