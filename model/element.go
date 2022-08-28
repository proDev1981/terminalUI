package model

import "strings"
import "strconv"

type Element struct{
	name string
	Style
	value string
	parent element
  Listeners
  prev *Element
}
// methods
func NewElement(name string, css Style, value string, listener Listeners)*Element{
  return &Element{name:name, Style:css, value:value, Listeners:listener } 
}

func (this *Element) Render(){
	DrawElement(&this.Style, this.value, this.Parent())
}

func (this *Element) Focus(){

  parent := this.Parent().(*box)
  if parent.prev != nil{
    parent.prev.Outside()
  }
  parent.prev = this
	MoveCursor(this.focus_x, this.focus_y)
  this.Inside()
  if this.Background != "" { Colored(this.Background) }
  if this.Color != "" { Colored(this.Color) }
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
  return this.IsEditable() || this.IsClickable() || this.IsBox()
}

func (this *Element) UploadStyles(types string){
  if types == "inside"{
    UploadStyles(&this.Style, this.Style.Inside)
  }
  if types == "outside"{
    UploadStyles(&this.Style, this.Style.Outside)
  }
}

func (this *Element) AddListener(types string, f func(e Event)){
  switch types{
  case "click":
    this.onClick = f
  case "inside":
    this.onInside = f
  case "outside":
    this.onOutside = f
  }
}

func (this *Element) Click(){
	this.onClick(Event{"click", this})
	this.Focus()
}
func (this *Element) OnClick(f func (e Event)){
  this.onClick = f
}

func (this *Element) Inside(){
  this.UploadStyles("inside")
  this.Render()
  this.onInside(Event{"inside", this})
}
func (this *Element) OnInside(f func (e Event)){
  this.onInside = f
}

func (this *Element) Outside(){
  this.UploadStyles("outside")
  this.Render()
  this.onOutside(Event{"outside", this})
}
func (this *Element) OnOutside(f func (e Event)){
  this.onOutside = f
}

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
