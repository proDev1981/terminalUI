package model

// box content elements
type box struct{
	Element
	childs []element
	focus int
}
func Box(name string, style Style, childs ...element)*box{
	return &box{Element{name, style,"", nil}, childs, 0 }
}
// methods inside interface element
func (b *box) Render(){
	// clean window
	Clean2()
	// draw element
	DrawElement(&b.Style, "", nil)
	// render childs
	for _,child := range b.childs{
		child.SetParent(b)
		child.Render()
	}
	b.childs[b.focus].Focus()

}

func (b *box) Focus(){
	b.childs[b.focus].Focus()
}

func (b *box) IsEditable()bool{
	return b.childs[b.focus].IsEditable()
}

func (b *box) IsClickable()bool{
	return b.childs[b.focus].IsClickable()
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
	if b.focus < len(b.childs)-1 { 
		b.focus++
		child := b.childs[b.focus]
		if child.IsEditable() || child.IsClickable(){
			b.childs[b.focus].Focus()
		}else{
			b.NextChild()
		}
	}else{
		b.focus = 0		
		child := b.childs[b.focus]
		if child.IsEditable() || child.IsClickable(){
			b.childs[b.focus].Focus()
		}else{
			b.NextChild()
		}
	}
}
