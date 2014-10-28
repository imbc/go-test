// https://github.com/ctessum/go.clipper

package main

type Path []*IntPoint

type PolyTree struct {
	PolyNode
	m_AllPolys []*PolyNode
}

type PolyNode struct {
	m_Parent   *PolyNode
	m_polygon  Path
	m_Index    int
	m_jointype JoinType
	m_endtype  EndType
	m_Childs   []*PolyNode
	IsOpen     bool
}

func NewPolyTree() *PolyTree {
	pt := new(PolyTree)
	pt.m_AllPolys = make([]*PolyNode, 0)
	return pt
}
