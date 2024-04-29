package practice

import . "AlgoPractice/src/graph/structure"

/*
Problem:
Ref: https://leetcode.com/problems/clone-graph/description/

Given:
- A reference of a node in connected undirected graph

Return:
- A deep copy of graph

Constraints:
- number of nodes: [1, 100]
- node value: [1, 100]
- node value is unique
- no repeated edges and no self-loops in the graph
- graph is connected and all nodes can be visited from the given node
*/

/*
First approach:
I will use DFS to copy the graph
*/

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[int]*Node)
	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}

		clone := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
		visited[node.Val] = clone
		for _, neighbor := range node.Neighbors {
			cloneNeighbor, exist := visited[neighbor.Val]
			if exist {
				clone.Neighbors = append(clone.Neighbors, cloneNeighbor)
			} else {
				clone.Neighbors = append(clone.Neighbors, dfs(neighbor))
			}
		}

		return clone
	}

	return dfs(node)
}
