package main


import m 	"framework/model"
import 		"fmt"



func main(){

	b := m.Box("box1", m.ClassBox,
			m.Label("label1", m.Class, "Password"),
			m.Button("btn1", m.Class, "aceptar", IncrementStatusBar),
			m.StatusBar("status-bar", "0", "%"),
			m.StatusBar("status-bar2", "0","$"),
	)
	b.Render()
	m.Loop(b, func(){})

}

func IncrementStatusBar(e m.Event){
	parent := e.Target.Parent()
	statusBar := parent.Select("status-bar").Select("bar-status-bar")
	porcent := parent.Select("status-bar").Select("porcent-status-bar")
	statusBar.SetStatus(statusBar.GetStatus()+1)
	porcent.SetValue(fmt.Sprint(statusBar.GetStatus()))
}



