package main

import (
	"cameraGT/cmd/custom"
	"cameraGT/cmd/model"
)

type CustomString2 []custom.IUltimate

func replace(s *[]string) {
	(*s)[0] = "999"
}

func main() {
	TestAny()
	//a := custom.StringUltimate{"1","2"}
	//a[0] = "100"
	//b:= []string(a)
	//replace(a.ToArray())
	//fmt.Println(b)
	//fmt.Println("waitting....")
	//fmt.Scanln()

	//a := app.New()
	//
	//w := a.NewWindow("Hello")
	//w.SetContent(widget.NewVBox(
	//	widget.NewLabel("Hello Fyne!"),
	//	widget.NewButton("Quit", func() {
	//		a.Quit()
	//	})))
	//
	//w.ShowAndRun()
}

func TestAny() {
	dog := model.Dog{Old: 15}
	cat := model.Cat{Name: "mèo"}
	inRa(dog)
	inRa(cat)
}

func inRa(any custom.IAny) {
	any.(model.Animal).Sua()
	switch ob := any.(type) {
	case model.Dog:
		ob.Sua()
	case model.Cat:
		ob.Sua()
	}
}
