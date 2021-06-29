package main

import (
	"fyne.io/fyne/app"

	"project6/view"
)

func main() {
	myApp := app.New()
	myWin := myApp.NewWindow("聊天室")
	myWin.SetContent(view.SetView())
	myWin.ShowAndRun()
}
