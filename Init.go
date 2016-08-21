package bl

import (
	"github.com/amortaza/go-g5"
)

func Init() {
	g5.Init()

	// initial resize fire
	resizeEvent := NewResizeEvent(g_hal.GetWindowDim())
	FireEvent(resizeEvent)
}

func Uninit() {
	g5.Uninit()
}

