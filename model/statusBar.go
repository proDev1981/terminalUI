package model

func StatusBar(name, initial, unit string)*box{

	return	Box(name,ClassStatusBar,
					Bar("bar-" + name, ClassBar),
					Label("porcent-" + name, ClassLabelBar, initial),
					Label("unit-" + name, ClassLabelBar, unit),
			)
		}
