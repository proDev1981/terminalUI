package model

import "strings"
import "fmt"
import "os"
import "strconv"
import k "github.com/eiannone/keyboard"

// functions 

func Clean(c, f int){
	MoveCursor(0, 0)
	Colored(B_BLACK)
	line := strings.Repeat(" ",f)
	print(line)
	row := 0
	for row < c {
		print(row)
		MoveCursor(0, row)
		print(line)
		row++
	}
}

func Clean2(){
	print("\033[2K")
}

func align(s *Style)int{
	var y int
	switch s.Align {
	case "center":
		position := s.Height/2
		y += s.Y + position
	case "top":
		y = s.Y
	case "botton":
		y = s.Y + s.Height-1
	default:	
		position := s.Height/2
		y = s.Y + position
	}
	return y
}

func justify(s *Style, value string)int{
	var x int
	switch s.Justify{
	case "center":
		x = s.X + (s.Width - len(value))/2
	case "rigth":
		x = s.X + s.Width - len(value) 
	case "left":
		x = s.X
	default:
		x = s.X + (s.Width - len(value))/2
	}
	return x
}

func SetWidth(n int, border []string){
	if len(border) > 0 {
		print(border[5]+ strings.Repeat(" ", n-2)+border[5])
	}else{
		print(strings.Repeat(" ", n ))
	}
}

func SetHeight(x, y, width, height int, border []string){
	for i := 0 ; i < height ; i++{
		MoveCursor(x, y+i)
		if len(border) > 0 && i == 0 {
			print(border[0]+ strings.Repeat(border[4], width-2)+border[1])
		}else if len(border) > 0 && i == height-1 {
			print(border[2]+ strings.Repeat(border[4], width-2)+border[3])

		}else{
			SetWidth(width, border)
		}
	}
}

func DrawElement(s *Style, value string, parent element){
	if s.Position == "" { s.Position = "static" } // manterner hasta que haga un objeto style standar
	isRelative(s, parent)
	isStatic(s, parent)
	MoveCursor(s.X, s.Y)
	if s.Background != "" { Colored(s.Background) }
	if len(value) > s.Width { s.Width = len(value)}
	SetHeight(s.X, s.Y, s.Width, s.Height, s.Border)
	s.focus_x = justify(s, value)
	s.focus_y = align(s)
	MoveCursor(s.focus_x, s.focus_y)
	if s.Color      != "" { Colored(s.Color) }
	print(value)
	ResetColors()
}

func isRelative(s *Style, parent element){
	if parent != nil {
		if s.Position == "relative" && !s.Built{
			p := parent.GetStyle()
			s.X += p.X
			s.Y += p.Y
			s.Built = true
		}
	}
}

func isStatic(s *Style, parent element){
	var s2 *Style
	if parent != nil {
		if s.Position == "static" && !s.Built {
			p := parent.GetStyle()
			if s.order != 0 {
				s2 = parent.(*box).childs[s.order-1].GetStyle()
			}else{
				s2 = &Style{}
			}
			if s.order == 0 {
				s.X += p.X + GetMarginLeft(s)
				s.Y += p.Y + GetMarginTop(s)
			}else{
				s.X += s2.X + s2.Width + GetMarginLeft(s) + GetMarginRight(s2)
				s.Y += p.Y + GetMarginLeft(s)
			}
			if s.X > p.Width + p.X {
				s.X = p.X + GetMarginLeft(s) + GetMarginRight(s2)
				s.Y = s2.Y + GetMarginTop(s) + s2.Height
			}
			s.Built = true
		}
	}
}

func ResetColors(){
	print("\033[0m")
}


func Loop(b *box, f func()){

	if err := k.Open(); err != nil { println(err) }
	defer k.Close()

	for{
		char, key, _ := k.GetKey()
		switch key {
			case k.KeyEsc:
				return
			case k.KeyTab:
				b.NextChild()
			case k.KeyEnter:
				b.Click()
			case k.KeySpace:
				if b.IsEditable(){
					if b.SetValue(" "){
						print(" ")
					}
				}
			case 127:
				if b.IsEditable(){
					value := b.GetValue()
					b.DelValue()
					if value != "" {
						b.SetValue(value[:len(value)-1])
						print("\033[D"," ", "\033[D")
					}
				}
			default:
				if b.IsEditable(){
					if b.SetValue(string(char)){
						print(string(char))
					}
				}
		}
		f()
	}
}
func Colored(color string){
	print(color)
}

func MoveCursor(line, raw int){
	print("\033["+fmt.Sprint(raw)+";"+fmt.Sprint(line)+"H")
}

func ColorRGB(r, g, b int)string{
	red := fmt.Sprint(r)
	green := fmt.Sprint(g)
	blue := fmt.Sprint(b)
	return "\033[38;2;"+red+";"+green+";"+blue+"m"
}

func BackRGB(r, g, b int)string{
	red := fmt.Sprint(r)
	green := fmt.Sprint(g)
	blue := fmt.Sprint(b)
	return "\033[48;2;"+red+";"+green+";"+blue+"m"
}
// debuggers
func debugger(value any)bool{
	file, err := os.OpenFile("debugger.log", os.O_RDWR,0755)
	if err != nil{
		file, err = os.Create("debugger.log")
		if err != nil{
			println("error de apertura: ",err)
			return false
		}
	}
	var c []byte
	_, err = file.Read(c)
	if err != nil{
		println("error de lectura:",err)
		return false
	}
	c = append(c, []byte(fmt.Sprintln(value))...)
	_, err = file.Write(c)
	if err != nil {
		fmt.Println("error escritura: ",err)
	}
	file.Close()
	return true
}

func IsFocusable(e element)bool{
	return 	e.IsEditable() || 
			e.IsClickable() || 
			e.IsBox()  
}
func GetMarginLeft(s *Style)(res int){
	strMargin := strings.Split(s.Margin, ",")
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
func GetMarginRight(s *Style)(res int){
	strMargin := strings.Split(s.Margin, ",")
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
func GetMarginBotton(s *Style)(res int){
	strMargin := strings.Split(s.Margin, ",")
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
func GetMarginTop(s *Style)(res int){
	strMargin := strings.Split(s.Margin, ",")
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
