package bl

import (
	"fmt"
	"container/list"
)

var Root_Node    *Node
var Current_Node *Node

var Window_Width, Window_Height int
var Mouse_X, Mouse_Y int

func Root() {

	Current_Node = newNode()

	Current_Node.Id = "ROOT"

	g_nodeById["ROOT"] = Current_Node

	g_nodeStack.Push(Current_Node)

	Root_Node = Current_Node
}

func Div() {
	parent := Current_Node

	Current_Node = newNode()

	parent.Kids.PushBack(Current_Node)
	Current_Node.Parent = parent

	g_nodeStack.Push(Current_Node)
}

func Id(id string) {
	Current_Node.Id = id
	g_nodeById[id] = Current_Node
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
	Current_Node.left = left
	Current_Node.top = top
}

func Dim(width, height int) {
	Current_Node.width = width
	Current_Node.height = height
}

func InvisibleToMouseEvents() {
	Current_Node.InvisibleToMouseEvents = true
}

func CustomRenderer(f func(node *Node), topsKids bool) {

	if Current_Node.CustomRender_1 == nil {
		Current_Node.CustomRender_1 = f

	} else {
		Current_Node.CustomRender_2 = f
	}

	Current_Node.CustomsShouldRendersAfterKids = topsKids
}

func Dirty() {
	Current_Node.Dirty = true
}

func SettleBoundary() {
	Current_Node.SettledBoundary = true
}

func SettleKids() {
	Current_Node.SettledKids = true
}

func RequireSettledBoundary()  {

	if !Current_Node.SettledBoundary {
		fmt.Println("Boundary has not been settled for node ", Current_Node.Id)
//		panic("See print out - RequireSettledBoundary error")
	}
}

func RequireSettledKids() {
	if !Current_Node.SettledKids {
		fmt.Println("Kids have not been settled for node ", Current_Node.Id)
		//panic("See print out - RequireSettledKids error")
	}
}

func OwnsLeft(owner string) bool {
	return Current_Node.OwnsLeft(owner)
}

func OwnsTop(owner string) bool {
	return Current_Node.OwnsTop(owner)
}

func OwnsWidth(owner string) bool {
	return Current_Node.OwnsWidth(owner)
}

func OwnsHeight(owner string) bool {
	return Current_Node.OwnsHeight(owner)
}

func OwnsPos(owner string) bool {
	return Current_Node.OwnsPos(owner)
}

func OwnsDim(owner string) bool {
	return Current_Node.OwnsDim(owner)
}

func OnMouseMove(cb func(*MouseMoveEvent)) {
	registerOnMouseMoveOnNode(Current_Node, cb)
}

func OnMouseButton(cb func(*MouseButtonEvent)) {
	registerOnMouseButtonOnNode(Current_Node, cb)
}

func AddStabilizeFunc_PreKids(cb func()) {

	if Current_Node.stabilize_funcs_pre_kids == nil {
		Current_Node.stabilize_funcs_pre_kids = list.New()
	}

	Current_Node.stabilize_funcs_pre_kids.PushBack(cb)
}

