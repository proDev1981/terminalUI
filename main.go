package main

import m "framework/model"
import "fmt"


func main(){

	b := m.Box("box1", m.Body,
                m.Label("Titulo", m.Tittle, "Add new item"),
                m.Input("name", m.InputText, "", " Name.."),
                m.Input("apellidos", m.InputText, "", " Apellidos.."),
                m.Input("edad", m.InputText, "", " Edad.."),
                m.Input("telefono", m.InputText, "", " Telefono.."),
			          m.Button("btn1", m.Btn, "Save", SaveData),
	)
	m.InitApp(b)
	m.Loop(b, func(){})

}
func SaveData(e m.Event){
  parent := e.Target.Parent()
  name := parent.Select("name").GetValue()
  apellidos := parent.Select("apellidos").GetValue()
  edad := parent.Select("edad").GetValue()
  telefono := parent.Select("telefono").GetValue()
  fmt.Println(name, apellidos, edad, telefono)
}



