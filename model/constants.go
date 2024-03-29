package model

var DEFAULT_LISTENERS = Listeners{ func(s Event){}, func(e Event){}, func(e Event){} }
var	BORDER_DEFAULT []string = []string{"┌","┐","└","┘","─","│"}
var BORDER_RADIUS []string = []string{"╭","╮","╰","╯","─","│"}

const(
	B_GREEN 		= "\033[42m"
	B_RED 			= "\033[41m"
	B_BLUE 			= "\033[44m"
	B_PURPLE 		= "\033[45m"
	B_WHITE 		= "\033[47m"
	B_BLACK 		= "\033[40m"
	B_GRAY 			= "\033[1;40m"
	B_BROWN 		= "\033[43m"
	GREEN 			= "\033[32m"
	RED 			= "\033[31m"
	BLUE 			= "\033[34m"
	PURPLE 			= "\033[35m"
	WHITE 			= "\033[37m"
	BLACK 			= "\033[30m"
	GRAY 			= "\033[1;30m"
	BROWN 			= "\033[33m"
	GREEN_LIGHT 	= "\033[1;32m"
	RED_LIGHT 		= "\033[1;31m"
	BLUE_LIGHT 		= "\033[1;34m"
	PURPLE_LIGHT 	= "\033[1;35m"
	WHITE_LIGHT 	= "\033[1;37m"
	BROWN_LIGHT 	= "\033[1;33m"
	TRANSPARENT 	= ""
	///////////////////

)


// interfaces
type element interface{
	Render()
	Focus()
	Name()string
	IsEditable()bool
	IsClickable()bool
	IsBox()bool
	IsFocusable()bool
	Click()
  Inside()
  Outside()
  OnClick(func (Event))
  OnInside(func(Event))
  OnOutside(func(Event))
	SetValue(string)bool
	GetValue()string
	DelValue()
	SetParent(element)
	Parent()element
	Select(string)element
	GetState()bool
	SetState()
	GetStyle()*Style
	SetStatus(int)
	GetStatus()int
	GetOrder()int
	SetOrder(int)
	GetMarginLeft()int
	GetMarginRight()int
	GetMarginBotton()int
	GetMarginTop()int
  UploadStyles(string)
  AddListener(types string, f func(e Event))

}

type Listeners struct{
  onClick func(Event)
  onInside func(Event)
  onOutside func(Event)
}

type Style struct{
	tag string
	Width, Height int
	Background, Color string
	X, Y int
	order int
	focus_x, focus_y int
	Margin string
	Border string
  BorderColor string
	Position string
	Align string
	Justify string
	Built bool
	Inside *Style
  Outside *Style
  Line int
  Playholder string
  PlayholderColor string
}

