package bl

var g_shadowNodeByID map[string] *ShadowNode

func EnsureShadow() (*ShadowNode) {

	shadow, ok := g_shadowNodeByID[Current_Node.Id]

	if !ok {

		shadow = NewShadowNode(Current_Node)

		g_shadowNodeByID[Current_Node.Id] = shadow
	}

	shadow._backingNode = Current_Node

	return shadow
}

func EnsureShadowById(id string) (*ShadowNode) {

	shadow, ok := g_shadowNodeByID[id]

	node := GetNodeById(id)

	if !ok {

		shadow = NewShadowNode(node)

		g_shadowNodeByID[id] = shadow
	}

	shadow._backingNode = node

	return shadow
}

func EnsureShadowByNode(node *Node) (*ShadowNode) {

	shadow, ok := g_shadowNodeByID[node.Id]

	if !ok {

		shadow = NewShadowNode(node)

		g_shadowNodeByID[node.Id] = shadow
	}

	shadow._backingNode = node

	return shadow
}
