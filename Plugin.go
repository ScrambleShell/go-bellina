package bl

import "container/list"

var g_pluginByName map[string] PlugIn
var g_pluginsInOrder *list.List

func init() {
	g_pluginByName = make(map[string] PlugIn)
	g_pluginsInOrder = list.New()
}

type PlugIn interface {
	Name() string
	Init()
	Tick()
	Reset()
	Uninit()
	OnNodeAdded(node *Node)
	OnNodeRemoved(node *Node)

	GetState() interface{}
	On(cb func(interface{}))
	On2(cb func(interface{}), start func(interface{}), end func(interface{}))
}

func Plugin(p PlugIn) {

	if _, has := g_pluginByName[p.Name()]; !has {
		p.Init()

		g_pluginByName[p.Name()] = p
		g_pluginsInOrder.PushBack(p)
	}
}

func GetPluginState(pluginName string) interface{} {
	plugin, ok := g_pluginByName[pluginName]

	if !ok {
		panic( "GetPluginState() cannot find unregistered plugin named " + pluginName)
	}

	return plugin.GetState()
}

func AddPluginOnTick(cb func()) {
	g_pluginTicks.PushBack(cb)
}
