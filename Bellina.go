package bl

import (
	"github.com/amortaza/go-g5"
	"container/list"
	"github.com/amortaza/go-bellina/core"
)

func Root() {
	Current_Node = NewNode()

	Current_Node.Id = "ROOT"

	g_nodeByID["ROOT"] = Current_Node

	g_nodeStack.Push(Current_Node)

	Root_Node = Current_Node
}

func DivId(id string) {
	g_nodeStack.Push(Current_Node)

	Current_Node = GetNodeById(id)
}

func Div() {
	parent := Current_Node

	Current_Node = NewNode()

	parent.Kids.PushBack(Current_Node)
	Current_Node.Parent = parent

	g_nodeStack.Push(Current_Node)
}

func Id(id string) {
	Current_Node.Id = id
	g_nodeByID[id] = Current_Node
}

func End() {
	g_nodeStack.Pop()

	if g_nodeStack.Size == 0 {
		Current_Node = nil
	} else {
		Current_Node = g_nodeStack.Top().(*Node)
	}
}

func Pos(left, top int) {
	Current_Node.Left, Current_Node.Top = left, top
}

func Dim(width, height int) {
	Current_Node.Width, Current_Node.Height = width, height
}

func Color(red,green,blue float32) {
	Current_Node.Red1, Current_Node.Green1, Current_Node.Blue1 = red, green, blue
}

func Colorv(color []float32) {
	Current_Node.Red1, Current_Node.Green1, Current_Node.Blue1 = color[0],color[1],color[2]
}

func Color2(red,green,blue float32) {
	Current_Node.Red2, Current_Node.Green2, Current_Node.Blue2 = red, green, blue
}

func FlagBorderAll() {
	Current_Node.Flags |= flag_BORDER_ALL
}

func FlagSkipRectDraw() {
	Current_Node.Flags |= flag_SKIP_NODE_RECT_DRAW
}

func FlagLabelCenterHorizontally() {
	Current_Node.Flags |= flag_LABEL_ALIGN_HCENTER
}

func FlagLabelCenterVertically() {
	Current_Node.Flags |= flag_LABEL_ALIGN_VCENTER
}

func Label(label string) {
	Current_Node.Label = label
}

func LabelOpacity(opacity float32) {
	Current_Node.LabelOpacity = opacity
}

func Font(fontName string, fontSize int) {
	Current_Node.FontName, Current_Node.FontSize = fontName, fontSize
}

func GetFont() (string, int) {
	return Current_Node.FontName, Current_Node.FontSize
}

func FontColor(red, green, blue float32) {
	Current_Node.FontRed, Current_Node.FontGreen, Current_Node.FontBlue = red, green, blue
}

func FontColorv(c []float32) {
	Current_Node.FontRed, Current_Node.FontGreen, Current_Node.FontBlue = c[0],c[1],c[2]
}

func FontNudge(x, y int) {
	Current_Node.FontNudgeX, Current_Node.FontNudgeY = x, y
}

func BorderThickness(thickness []int) {
	Current_Node.BorderThickness = thickness
}

func BorderColor(red, green, blue float32) {
	Current_Node.BorderRed, Current_Node.BorderGreen, Current_Node.BorderBlue = red, green, blue
}

func BorderTopsCanvas() {
	Current_Node.BorderTopsCanvas = true
}

func NodeOpacity1f(opacity float32) {
	Current_Node.NodeOpacity = []float32{opacity,opacity,opacity,opacity}
}

func NodeOpacity4f(opacity []float32) {
	Current_Node.NodeOpacity = opacity
}

func Texture(partialname string) {
 	texture, ok := g_textureByPartialName[partialname]

	if !ok {
		texture = g5.NewTexture()
		texture.LoadImage("github.com/amortaza/go-bellina-examples/assets/images/" + partialname + ".png")

		g_textureByPartialName[partialname] = texture
	}

	Current_Node.Texture = texture

	Current_Node.Width = texture.Width
	Current_Node.Height = texture.Height

	Current_Node.SeeThru = true
}

func PreventBubbling() {
	Current_Node.PreventBubbling = true
}

func InvisibleToEvents() {
	Current_Node.InvisibleToEvents = true
}

func OnMouseMove(cb func(*MouseMoveEvent)) {
	OnMouseMoveOnNode(Current_Node, cb)
}

func OnMouseMoveOnNode(node *Node, cb func(*MouseMoveEvent)) {
	if node.OnMouseMoveCallbacks == nil {
		node.OnMouseMoveCallbacks = list.New()
	}

	node.OnMouseMoveCallbacks.PushBack(cb);
}

func OnMouseButton(cb func(*MouseButtonEvent)) {
	OnMouseButtonOnNode(Current_Node, cb)
}

func OnMouseButtonOnNode(node *Node, cb func(*MouseButtonEvent)) {
	if node.OnMouseButtonCallbacks == nil {
		node.OnMouseButtonCallbacks = list.New()
	}

	node.OnMouseButtonCallbacks.PushBack(cb);
}

func GetNodeById(id string ) *Node {
	node, _ := g_nodeByID[id]

	return node
}

func GetFontHeight() int {

	fontname, fontsize := Current_Node.FontName, Current_Node.FontSize

	g4font := core.GetG4Font(fontname, fontsize)

	return g4font.Height
}

func CustomRenderer1(f func(node *Node), topsLabel bool) {
	Current_Node.CustomRender1_Hook = f
	Current_Node.CustomRender1_TopsLabel = topsLabel
}

