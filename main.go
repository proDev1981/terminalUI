package main


//import k "github.com/eiannone/keyboard"
import m "framework/model"
import "fmt"

//palette



// styles
var classBox m.Style = m.Style{

	X 			: 20,
	Y 			: 20,
	Background 	: m.BackRGB(250,250,250),
	Width 		: 112,
	Height		: 5,

}

var classLabel m.Style = m.Style{

	X 			: 2,
	Y 			: 1,
	Width		: 20,
	Height		: 3,
	Color		: m.BLACK,
	Background	: m.BackRGB(250,250,250),

}

var classInput m.Style = m.Style{

	X 			: 22,
	Y 			: 2,
	Width		: 20,
	Height		: 1,
	Background	: m.B_BLACK,
	Color		: m.BackRGB(250,250,250),
	Justify 	: "left",

}

var classAceptar m.Style = m.Style{
	X 			: 44,
	Y 			: 1,
	Width 		: 20,
	Height 		: 3,
	Color		: m.BLACK,
	Background	: m.BackRGB(250,250,250),
	Border 		: m.BORDER_DEFAULT,
	Align 		: "center",
	Justify 	: "center",
}

var classToggle m.Style = m.Style{
	X 			: 64,
	Y 			: 1,
	Width 		: 20,
	Height 		: 3,
	Color		: m.BLACK,
	Background	: m.BackRGB(250,250,250),
	Align 		: "center",
	Justify 	: "center",
}

var classBar m.Style = m.Style{
	X 			: 84,
	Y 			: 2,
	Width 		: 20,
	Height 		: 1,
	Color		: m.BackRGB(10,250,10),
	Background	: m.BackRGB(10,100,10),
	Align 		: "center",
	Justify 	: "left",
}

var classLabelBar m.Style = m.Style{

	X 			: 104,
	Y 			: 2,
	Width		: 5,
	Height		: 1,
	Color		: m.BLACK,
	Background	: m.BackRGB(250,250,250),

}


func main(){

	var buf string
	b := m.Box(
			"box1",
			classBox,
			m.Label("label1", classLabel, "Password"),
			m.Input("input1", classInput, buf),
			m.Button("Aceptar", classAceptar, "Aceptar", func(e m.Event){
				bar := e.Target.Parent().Select("bar")
				labelBar := e.Target.Parent().Select("labelBar")
				bar.SetStatus(bar.GetStatus()+1)
				labelBar.SetValue(fmt.Sprint(" ",bar.GetStatus(),"%"))
			}),
			m.Toggle("toggle1",classToggle, " toggle1", func(e m.Event){
			}),
			m.StatusBar("bar", classBar),
			m.Label("labelBar",classLabelBar, " 0%"),
	)

	b.Render()
	m.Loop(b, func(){})

}





