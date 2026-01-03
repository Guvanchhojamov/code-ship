package graphs

// func main() {
// 	fmt.Println("Hello, World!")
// 	ds := NewDisjointSet(5)
// 	fmt.Println(ds.findParent(0), ds.parents, ds.ranks)
// 	fmt.Println(ds.findParent(1), ds.parents, ds.ranks)
// 	ds.unionByRank(0, 1)
// 	fmt.Println(ds.findParent(0), ds.parents, ds.ranks)
// 	fmt.Println(ds.findParent(1), ds.parents, ds.ranks)
// 	//
// 	fmt.Println(ds.findParent(2), ds.parents, ds.ranks)
// 	fmt.Println(ds.findParent(3), ds.parents, ds.ranks)
// 	ds.unionByRank(2, 3)
// 	fmt.Println(ds.findParent(2), ds.parents, ds.ranks)
// 	fmt.Println(ds.findParent(3), ds.parents, ds.ranks)

// 	fmt.Println(ds.findParent(0), ds.parents, ds.ranks)
// 	fmt.Println(ds.findParent(3), ds.parents, ds.ranks)
// 	ds.unionByRank(0, 3)
// 	fmt.Println(ds.findParent(0), ds.parents, ds.ranks)
// 	fmt.Println(ds.findParent(3), ds.parents, ds.ranks)
// 	//
// 	fmt.Println(ds.findParent(1), ds.parents, ds.ranks)
// 	fmt.Println(ds.findParent(2), ds.parents, ds.ranks, ds.sizes)
// 	ds.unionByRank(1, 2)
// 	fmt.Println(ds.parents, ds.ranks, ds.sizes)
// }

type DisjointSet struct {
	parents, sizes, ranks []int
}

func NewDisjointSet(n int) DisjointSet {
	ds := DisjointSet{
		parents: make([]int, n),
		sizes:   make([]int, n),
		ranks:   make([]int, n),
	}
	for i := range ds.parents {
		ds.parents[i] = i
		ds.sizes[i] = 1
	}
	return ds
}

func (d *DisjointSet) unionBySize(u, v int) {
	// when we mege sets by size, we need to merge small size to bigger size.
	// why? because this optimizes find() in flat graphs.
	up := d.findParent(u)
	vp := d.findParent(v)

	if up == vp {
		return
	}
	if d.sizes[vp] > d.sizes[up] {
		d.parents[up] = vp
		d.sizes[vp] += d.sizes[up]
	} else {
		d.parents[vp] = up
		d.sizes[up] += d.sizes[vp]
	}
}

func (d *DisjointSet) unionByRank(u, v int) {
	// we need to merge(union) nodes only it they are not in same set.
	// how we know this? Taking ultimate parents of u,v nodes.
	// if they ult parents are same they aready merged.
	u_p := d.findParent(u)
	v_p := d.findParent(v)
	if u_p == v_p {
		return
	}
	if d.ranks[v_p] > d.ranks[u_p] {
		d.parents[u_p] = v_p
	} else if d.ranks[u_p] > d.ranks[v_p] {
		d.parents[v_p] = u_p
	} else {
		d.parents[v_p] = u_p
		d.ranks[u_p] += 1 // increment only when equal rank
	}
}

// find ultimate parent for u,v nodes. to chek is they in same group or same path.
func (d *DisjointSet) findParent(node int) int {
	//TODO find ultimate parent for node.
	if d.parents[node] == node {
		return node
	}
	ult_p := d.findParent(d.parents[node])
	d.parents[node] = ult_p // optimize with compress node path to access in O(1) in next time.
	return ult_p
}

/*
Disjoint Set

Design a disjoint set (also called union-find)
data structure that supports the following operations:

DisjointSet(int n) initializes the disjoint set with n elements.
    void unionByRank(int u, int v) merges the sets containing u and v using the rank heuristic.
    void unionBySize(int u, int v) merges the sets containing u and v using the size heuristic.
    bool find(int u, int v) checks if the elements u and v are in the same set and returns true if they are, otherwise false.
        find() means check u and v are in same set ,how we can check it?
        if the are under the same paretn they are in same set, or in same component graph.

Example 1
Input:
["DisjointSet", "unionByRank", "unionBySize", "find", "find"]
[[5], [0, 1], [2, 3], [0, 1], [0, 3]]

Output:
[null, null, null, true, false]


*/
/*
Disjoint set functions:
    unionBySize, unionByRank, find OR findUltimateParent.
How works disjoint set?
1 we need to implement n elements, nodes.
 in init:
    - create parent arr, with len n.
    - i - is node, arr[i] is its parent.
    - create sizes array, with size n.
    - create ranks array with size n.
    - fill sizes and ranks array with 0.
at first it means
we have n nodes seperated.
all nodes has same rank/size.
Now we need to connect UnionByRank
nodes u,v.
    rules:
   if ranks[u] > ranks[v]:
     // we set parent of [v] as [u] in
        parents arr. and ranks[u]+1
     parents[v] = u
     ranks[u]=+1
   else:
    parents[u] = v
    ranks[v] = u
   // like this we union by rank.how. with size we see later.
Find(u,v) bool:
    in this func we need to find ultimate parent of both,
    if they are same return true otherwise return false.
while parents[u] != u:
    u = parents[u]
while parents[v] != v:
    v = parents[v]
if u == v:
    return False
else:
    return True

So like this we can union 2 nodes in O(1) time
and find() also in O(1) time.
    But find is works diff, its not exact O(1),
    its near about ~O(1).
But we use path compression algo in this function.
    we set ultimate parent after first time if its not ultimate
parent for node..
    for find this in next time, find in O(1).
    without changing rating.
Where we use this technick?
  Where the graph is dynamically changes, and added some node, deleted some node.
   and we need to check in O(1) time 2 node ult parents.
or add or delete in O(1) time..
 Notes:
    in DSU we prioritize only root nodes.
    child nodes does not matters. they just help us to find root for each subset...
    so we ignore additional operations for child nodes, only use path compression..
*/
