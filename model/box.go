package model


// box content elements
type box struct{
	Element
	childs []element
	focus int
}
func Box(name string, style Style, childs ...element)*box{
  return &box{Element{name:name, Style:style, Listeners:DEFAULT_LISTENERS }, childs, 0 }
}
// methods inside interface element
func (b *box) Render(){
	// clean window
	// draw element
	DrawElement(&b.Style, "", b.Parent())
	// render childs
	for index, child := range b.childs{
		child.SetParent(b)
		child.SetOrder(index)
		child.Render()
	}
	b.childs[b.focus].Focus()

}

func (b *box) Focus(){
	child := b.childs[b.focus]
	if IsFocusable(child){ 
		child.Focus() 
	}else{
		b.NextChild()
	}
}

func (b *box) IsEditable()bool{
	return b.childs[b.focus].IsEditable()
}

func (b *box) IsClickable()bool{
	return b.childs[b.focus].IsClickable()
}

func (b *box) IsBox()bool{
	return true
}

func (b *box) Click(){
	if b.IsClickable(){
		b.childs[b.focus].Click()
	}
}

func (b *box) SetValue(v string)bool{
	return b.childs[b.focus].SetValue(v)
	
}

func (b *box) GetValue()string{
	return b.childs[b.focus].GetValue()
}

func (b *box) DelValue(){
	b.childs[b.focus].DelValue()
}
// methods outside interface element
func (b *box) Select(name string)element{
	for _, child := range b.childs{
		if child.Name() == name {
			return child
		}
	}
	return nil
}


func (b *box) NextChild(){

	switch {
	case b.InRange() :
		b.focus++

	case b.childs[b.focus].IsBox() :
		  if b.childs[b.focus].(*box).InLastSon() {
			  b.childs[b.focus].(*box).focus = 0
			  if b.InRange(){ b.focus++ } else { b.focus = 0 }
		  }else if b.childs[b.focus].(*box).InRange(){
			  b.childs[b.focus].(*box).NextChild()
		  }

	default:
    if !b.ChildsContainsFocusables(){
      parent := b.Parent()
      if parent.(*box).InRange(){ parent.(*box).focus++ } else { parent.(*box).focus = 0 }
      parent.Focus()
      return
    }
		b.focus = 0
	}
	b.Focus()
}

func (b *box) InLastSon()bool{
	return b.focus == len(b.childs)-1
}

func (b *box) InRange()bool{
	return b.focus < len(b.childs)-1
}

func (b *box) ChildsContainsFocusables()bool{
  focusable := false
  for _, item := range b.childs{
    if item.IsEditable() || item.IsClickable(){
      focusable = true
    }
  }
  return focusable
}

func (b *box) GetChilds()[]element{
	return b.childs
}

