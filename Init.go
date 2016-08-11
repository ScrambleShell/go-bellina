package bl

import (
	"github.com/amortaza/go-g5"
	"github.com/amortaza/go-bellina/event"
	"container/list"
)

func Init() {
	g5.Init(/*DevicePixelRatio*/)

	g_shadowNodeByID = make(map[string] *ShadowNode)
	g_textureByPartialName = make(map[string] *g5.Texture)

	g_pluginTicks = list.New()

	// initial resize fire
	resizeEvent := event.NewResizeEvent(g_hal.GetWindowDim())
	event.Fire(resizeEvent)
}

func Uninit() {
	// free plugins
	for _, plugin := range g_pluginByName {

		plugin.Uninit()
	}

	// free texture map images
	for _, value := range g_textureByPartialName {
		value.Free()
	}

	g5.Uninit()
}

