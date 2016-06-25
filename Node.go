package bl

import (
	"container/list"
<<<<<<< HEAD
	"github.com/amortaza/go-g4"
=======
	"g4"
>>>>>>> 457139a89938fb2b2ae6a239356e9473e6b022ec
)

type Node struct {
	ID string

	Left, Top, Width, Height int32

	Red1, Green1, Blue1 float32
	Red2, Green2, Blue2 float32

	NodeOpacity []float32

	Flags uint32

	Label string
	LabelOpacity float32

	FontName string
	FontSize int32
	FontRed, FontGreen, FontBlue float32
	FontNudgeX, FontNudgeY int32

	BorderThickness []int32
	BorderRed, BorderGreen, BorderBlue float32
	BorderTopsCanvas bool

	// no need to free this - this is globally managed
	Texture *g4.Texture

	SeeThru bool

	Parent *Node
	Kids *list.List

	OnMouseMoveCallbacks *list.List
	OnMouseButtonCallbacks *list.List
}

func NewNode() *Node {
	node := &Node{}

	node.Flags = Z_COLOR_SOLID
	node.LabelOpacity = 1
	node.NodeOpacity = FourOnesFloat

	node.Kids = list.New()
	
	return node
}

func (node *Node) CallMouseMoveCallbacks(e *MouseMoveEvent) {
	if node.OnMouseMoveCallbacks != nil {
		for element := node.OnMouseMoveCallbacks.Front(); element != nil; element = element.Next() {
			cb := element.Value.(func(*MouseMoveEvent))

			cb(e)
		}
	}
}

func (node *Node) CallMouseButtonCallbacks(e *MouseButtonEvent) {
	if node.OnMouseButtonCallbacks != nil {
		for element := node.OnMouseButtonCallbacks.Front(); element != nil; element = element.Next() {
			cb := element.Value.(func(*MouseButtonEvent))

			cb(e)
		}
	}
}

func (node *Node) Free() {
	// free shadow node here!
}