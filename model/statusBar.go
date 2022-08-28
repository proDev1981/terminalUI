package model

func StatusBar(name, initial, unit string)*box{

	return	Box(name,ClassStatusBar,
					Bar("bar-" + name, ClassBar),
					Label("porcent-" + name, ClassLabelBar, initial),
					Label("unit-" + name, ClassLabelBar, unit),
			)
		}

  //styles



var ClassStatusBar Style = Style{

	Width		    : 45,
	Height		  : 5,
	Margin 		  : "0,0,0,0",
	Color		    : BLACK,
	Background	: BackRGB(250,250,250),
}

var ClassBar Style = Style{

	Width 		  : 20,
	Height 		  : 1,
	Margin 		  : "1",
	Justify 	  : "left",
	Color 		  : BackRGB(100, 20, 34),
	Background 	: BackRGB(0, 150, 0),
}

var ClassLabelBar Style = Style{

	Width 		  : 4,
	Height 		  : 1,
	Margin 		  : "1",
	Color 		  : BLACK,
	Background 	: ClassStatusBar.Background,
}

