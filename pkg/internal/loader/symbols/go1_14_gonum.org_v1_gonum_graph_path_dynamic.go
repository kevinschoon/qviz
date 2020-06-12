// Code generated by 'goexports gonum.org/v1/gonum/graph/path/dynamic'. DO NOT EDIT.

// +build go1.14,!go1.15

package symbols

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path/dynamic"
	"reflect"
)

func init() {
	Symbols["gonum.org/v1/gonum/graph/path/dynamic"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewDStarLite": reflect.ValueOf(dynamic.NewDStarLite),

		// type definitions
		"DStarLite":  reflect.ValueOf((*dynamic.DStarLite)(nil)),
		"WorldModel": reflect.ValueOf((*dynamic.WorldModel)(nil)),

		// interface wrapper definitions
		"_WorldModel": reflect.ValueOf((*_gonum_org_v1_gonum_graph_path_dynamic_WorldModel)(nil)),
	}
}

// _gonum_org_v1_gonum_graph_path_dynamic_WorldModel is an interface wrapper for WorldModel type
type _gonum_org_v1_gonum_graph_path_dynamic_WorldModel struct {
	WAddNode         func(a0 graph.Node)
	WEdge            func(uid int64, vid int64) graph.Edge
	WFrom            func(id int64) graph.Nodes
	WHasEdgeBetween  func(xid int64, yid int64) bool
	WHasEdgeFromTo   func(uid int64, vid int64) bool
	WNewNode         func() graph.Node
	WNewWeightedEdge func(from graph.Node, to graph.Node, weight float64) graph.WeightedEdge
	WNode            func(id int64) graph.Node
	WNodes           func() graph.Nodes
	WSetWeightedEdge func(e graph.WeightedEdge)
	WTo              func(id int64) graph.Nodes
	WWeight          func(xid int64, yid int64) (w float64, ok bool)
	WWeightedEdge    func(uid int64, vid int64) graph.WeightedEdge
}

func (W _gonum_org_v1_gonum_graph_path_dynamic_WorldModel) AddNode(a0 graph.Node) { W.WAddNode(a0) }
func (W _gonum_org_v1_gonum_graph_path_dynamic_WorldModel) Edge(uid int64, vid int64) graph.Edge {
	return W.WEdge(uid, vid)
}
func (W _gonum_org_v1_gonum_graph_path_dynamic_WorldModel) From(id int64) graph.Nodes {
	return W.WFrom(id)
}
func (W _gonum_org_v1_gonum_graph_path_dynamic_WorldModel) HasEdgeBetween(xid int64, yid int64) bool {
	return W.WHasEdgeBetween(xid, yid)
}
func (W _gonum_org_v1_gonum_graph_path_dynamic_WorldModel) HasEdgeFromTo(uid int64, vid int64) bool {
	return W.WHasEdgeFromTo(uid, vid)
}
func (W _gonum_org_v1_gonum_graph_path_dynamic_WorldModel) NewNode() graph.Node { return W.WNewNode() }
func (W _gonum_org_v1_gonum_graph_path_dynamic_WorldModel) NewWeightedEdge(from graph.Node, to graph.Node, weight float64) graph.WeightedEdge {
	return W.WNewWeightedEdge(from, to, weight)
}
func (W _gonum_org_v1_gonum_graph_path_dynamic_WorldModel) Node(id int64) graph.Node {
	return W.WNode(id)
}
func (W _gonum_org_v1_gonum_graph_path_dynamic_WorldModel) Nodes() graph.Nodes { return W.WNodes() }
func (W _gonum_org_v1_gonum_graph_path_dynamic_WorldModel) SetWeightedEdge(e graph.WeightedEdge) {
	W.WSetWeightedEdge(e)
}
func (W _gonum_org_v1_gonum_graph_path_dynamic_WorldModel) To(id int64) graph.Nodes { return W.WTo(id) }
func (W _gonum_org_v1_gonum_graph_path_dynamic_WorldModel) Weight(xid int64, yid int64) (w float64, ok bool) {
	return W.WWeight(xid, yid)
}
func (W _gonum_org_v1_gonum_graph_path_dynamic_WorldModel) WeightedEdge(uid int64, vid int64) graph.WeightedEdge {
	return W.WWeightedEdge(uid, vid)
}
