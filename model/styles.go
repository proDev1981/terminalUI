package model

// styles
var Body Style = Style{
	
	X 			    : 10,
	Y 			    : 3,
	Background 	: BackRGB(250,250,250),
	Width 		  : X100(50, TWidth()),
	Height		  : 50,
	Margin 		  : "0,0,0,0",
	Position 	  : "absolute",
}
var Btn Style = Style{

	Width		    : 20,
	Height		  : 3,
	Margin 		  : "1,1,1,1",
  Align       : "center",
  Border      : "radius",
  BorderColor : ColorRGB(200, 200, 200),
	Color		    : ColorRGB(0, 0, 0),
	Background	: Body.Background,
  Inside      : &Style{ BorderColor: ColorRGB(0, 0, 0) ,},
  Outside     : &Style{ BorderColor: ColorRGB(200, 200, 200),},
}

var Tittle Style = Style{

	Width		    : X100(90, Body.Width),
	Height		  : 1,
	Margin 		  : "1,20,2,2",
  Justify     : "center",
  Align       : "center",
	Color		    : BLACK,
	Background	: Body.Background,
}

var InputText Style = Style{

	Width		    : 20,
	Height		  : 1,
	Margin 		  : "1,5,2,2",
	Justify 	  : "left",
  Align       : "center",
	Color		    : ColorRGB(0, 0, 0),
	Background	: BackRGB(200, 200, 200),
  PlayholderColor : ColorRGB(100, 100, 100),
  Inside      : &Style{ Background : BackRGB(180, 180, 180), PlayholderColor: ColorRGB(180, 180, 180),},
  Outside     : &Style{ Background : BackRGB(200, 200, 200), PlayholderColor: ColorRGB(100, 100, 100),},
}

