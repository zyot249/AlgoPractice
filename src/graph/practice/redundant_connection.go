package practice

/*
Problem:
Ref: https://leetcode.com/problems/redundant-connection/description/

Given:
- An graph that started with n nodes labeled from 1 -> n, with 1 additional edge added
- The graph is represented as an array (edges) where edges[i] = [ai, bi]

Return:
- Return an edge that can be removed so that the resulting graph is a tree of n nodes.
- If there are multiple answers, return the answer that occurs last in the input.

Constraints:
- edges length == n
- n: [3, 1000]
- graph is connected
*/

/*
First approach:
A connected graph of n nodes and n edges has only 1 cycle.
From input, remove all the edges of the nodes that have only 1 edge
The rest will be the edges of cycle, choose the last
*/

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	for {
		edgeCount := make([]int, n+1)
		for _, edge := range edges {
			edgeCount[edge[0]]++
			edgeCount[edge[1]]++
		}

		var edgesAfter [][]int
		for _, edge := range edges {
			if edgeCount[edge[0]] > 1 && edgeCount[edge[1]] > 1 {
				edgesAfter = append(edgesAfter, edge)
			}
		}

		if len(edgesAfter) == len(edges) {
			break
		} else {
			edges = edgesAfter
		}
	}

	return edges[len(edges)-1]
}
