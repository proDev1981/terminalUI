package model

import "strings"
import "fmt"
import "os"
import "os/exec"
import "strconv"
import k "github.com/eiannone/keyboard"
import "assets"

// functions 
func Clear(){
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  err := cmd.Run()
  assets.Throw(err, "", nil)
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
  var border []string
	if s.Position == "" { s.Position = "static" } // manterner hasta que haga un objeto style standar
	isRelative(s, parent)
	isStatic(s, parent)
	MoveCursor(s.X, s.Y)
	if s.Background != "" { Colored(s.Background) }
	if s.Color      != "" { Colored(s.Color) }
	if len(value) > s.Width { s.Width = len(value)}
  if s.Border == "radius" { border = BORDER_RADIUS }
  if s.Border == "default" { border = BORDER_DEFAULT }
  if s.BorderColor != "" { Colored(s.BorderColor) }
	SetHeight(s.X, s.Y, s.Width, s.Height, border)
	s.focus_x = justify(s, value)
	s.focus_y = align(s)
	MoveCursor(s.focus_x, s.focus_y)
  if value == ""{
    Colored(s.PlayholderColor)
    print(s.Playholder)
  }
  Colored(s.Color)
	print(value)
	ResetColors()
  //print(s.X,":",s.Y)
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
  line := 0
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
        if parent.GetStyle().Line != 0 {
          line = parent.GetStyle().Line
        }else{ line = p.Y }
				s.X += s2.X + s2.Width + GetMarginLeft(s) + GetMarginRight(s2)
				s.Y += line + GetMarginTop(s)
			}
      if s.X + s.Width > p.Width + p.X {
        parent.GetStyle().Line = GetMaxHeight(parent)
        s.X = p.X + GetMarginLeft(s)
        s.Y = parent.GetStyle().Line + GetMarginTop(s)
      }
			s.Built = true
		}
	}
}
func GetMaxHeight(parent element)int{
  maxHeight := 0
  for _,item := range parent.(*box).childs{
    
    totalHeight := item.GetStyle().Y + item.GetStyle().Height 
    if totalHeight > maxHeight { maxHeight = totalHeight}
  }
  return maxHeight
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
        ResetColors()
        //Clear()
        Exit()
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

func GetSizeTerminal()(res []int){
  cmd := exec.Command("stty", "size")
  cmd.Stdin = os.Stdin
  out, err := cmd.Output()
  assets.Throw(err, "", nil)
  splitOut := strings.Split(string(out), " ")
  for _, item := range splitOut{
    item = strings.Trim(item, "\n ")
    num , err := strconv.Atoi(item)
    assets.Throw(err, "error al convertir string a intteger.", nil)
    res = append(res, num)
  }
  return res
}
func UploadStyles(ori , dest *Style){
  if dest != nil{
    if dest.Width  != 0{ ori.Width  = dest.Width }
    if dest.Height != 0{ ori.Height = dest.Height }
    if dest.Background  != "" { ori.Background  = dest.Background }
    if dest.Color  != "" { ori.Color  = dest.Color }
    if dest.Margin  != "" { ori.Margin  = dest.Margin }
    if dest.Border  != "" { ori.Border  = dest.Border }
    if dest.Position  != "" { ori.Position  = dest.Position }
    if dest.Align  != "" { ori.Align  = dest.Align }
    if dest.Justify  != "" { ori.Justify  = dest.Justify }
    if dest.Playholder  != "" { ori.Playholder  = dest.Playholder }
    if dest.PlayholderColor  != "" { ori.PlayholderColor  = dest.PlayholderColor }
    if dest.BorderColor != "" { ori.BorderColor = dest.BorderColor }
  }
}

func X100(porcent, parent int)int{
  return parent*porcent/100
}

func TWidth()int{
  return GetSizeTerminal()[1]
}

func THeight()int{
  return GetSizeTerminal()[0]
}

func Enter(){
  fmt.Println("\033[?1049h")
}

func Exit(){
  fmt.Println("\033[?1049l")
}

func InitApp(e element){
  Enter()
  e.Render()
}
