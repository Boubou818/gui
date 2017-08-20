package win

import (
	"github.com/faiface/gui/event"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func (w *Win) setUpMainthreadEvents() {
	var moX, moY int

	w.w.SetMouseButtonCallback(func(_ *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
		switch action {
		case glfw.Press:
			w.mainthreadEvent("mo", "down", moX, moY)
		case glfw.Release:
			w.mainthreadEvent("mo", "up", moX, moY)
		}
	})

	w.w.SetCursorPosCallback(func(_ *glfw.Window, x, y float64) {
		moX, moY = int(x), int(y)
		w.mainthreadEvent("mo", "move", moX, moY)
	})

	w.w.SetCharCallback(func(_ *glfw.Window, r rune) {
		w.mainthreadEvent("kb", "type", r)
	})

	w.w.SetSizeCallback(func(_ *glfw.Window, width, height int) {
		w.resize(width, height)
		w.mainthreadEvent("resize", 0, 0, width, height)
	})

	w.w.SetCloseCallback(func(_ *glfw.Window) {
		w.mainthreadEvent("wi", "close")
	})
}

func (w *Win) mainthreadEvent(a ...interface{}) {
	go func() {
		w.mainthreadEvents <- event.Sprint(a...)
	}()
}