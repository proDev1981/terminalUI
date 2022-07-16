package model

// styles
var ClassBox Style = Style{
	
	X 			: 10,
	Y 			: 10,
	Background 	: BackRGB(250,250,250),
	Width 		: 112,
	Height		: 5,
	Margin 		: "0,0,0,0",
	Position 	: "absolute",
}

var Class Style = Style{

	Width		: 20,
	Height		: 3,
	Margin 		: "1,1,1,1",
	Color		: BLACK,
	Background	: BackRGB(250,50,50),
}

var ClassInput Style = Style{

	Width		: 20,
	Height		: 1,
	Margin 		: "1,5,2,2",
	Justify 	: "left",
	Color		: B_WHITE,
	Background	: BackRGB(0, 0, 0),
}

var ClassStatusBar Style = Style{

	Width		: 45,
	Height		: 5,
	Margin 		: "0,0,0,0",
	Color		: BLACK,
	Background	: BackRGB(250,250,250),
}

var ClassBar Style = Style{

	Width 		: 20,
	Height 		: 1,
	Margin 		: "1,1,2,2",
	Justify 	: "left",
	Color 		: BackRGB(100, 20, 34),
	Background 	: BackRGB(0, 150, 0),
}

var ClassLabelBar Style = Style{

	Width 		: 4,
	Height 		: 1,
	Margin 		: "2,2,3,3",
	Color 		: BLACK,
	Background 	: ClassStatusBar.Background,
}


