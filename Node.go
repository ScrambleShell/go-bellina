package bl

import (
	"container/list"
	"github.com/amortaza/go-g4"
)

type Node struct {
	Id                                 string

	Left, Top, Width, Height           int

	Red1, Green1, Blue1                float32
	Red2, Green2, Blue2                float32

	NodeOpacity                        []float32

	Flags                              uint32

	Label                              string
	LabelOpacity                       float32

	FontName                           string
	FontSize                           int
	FontRed, FontGreen, FontBlue       float32
	FontNudgeX, FontNudgeY             int

	BorderThickness                    []int
	BorderRed, BorderGreen, BorderBlue float32
	BorderTopsCanvas                   bool

	// no need to free this - this is globally managed
	Texture                            *g4.Texture

	SeeThru                            bool

	Parent                             *Node
	Kids                               *list.List

	OnMouseMoveCallbacks               *list.List
	OnMouseButtonCallbacks             *list.List

	CustomRender1_Hook		   func()
	CustomRender1_TopsLabel            bool

	PreventBubbling bool
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