package main

type Dialog interface {
	createButton() interface{}
	render()
	onClick()
}

type WindowsDialog struct {
	Dialog
}
func (*WindowsDialog) createButton() WindowsButton {
	return WindowsButton{}
}

type HTMLDialog struct {
	Dialog
}
func (*HTMLDialog) createButton() HTMLButton {
	return HTMLButton{}
}

// 产品接口中将声明所有具体产品都必须实现的操作。
type Button struct {}
func (*Button) render(a, b int) {}
func (*Button) onClick(f int) {}

// 具体产品需提供产品接口的各种实现。
type WindowsButton struct {}
func (*WindowsButton) render(a, b int) {}
func (*WindowsButton) onClick(f int) {}

type HTMLButton struct {}
func (*HTMLButton) render(a, b int) {}
func (*HTMLButton) onClick(f int) {}

//func main() {
//	var h HTMLDialog
//	htmlButton := h.createButton()
//	h.Dialog.createButton()
//}
