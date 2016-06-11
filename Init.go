package bl

import (
	"g4"
	"bellina/event"
	"xel"
)

func Init() {
	g4.Init()

	g_shadowNodeByID = make(map[string] *ShadowNode)
	g_textureByPartialName = make(map[string] *g4.Texture)

	// initial resize fire
	resizeEvent := event.NewResizeEvent(xel.WinWidth, xel.WinHeight)
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

	g4.Uninit()
}

