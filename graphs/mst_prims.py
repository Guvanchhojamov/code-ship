class Solution:
    def spanningTree(self, V, edges):
        graph = [[] for _ in range(V)]
        for u, v, w in edges:
            graph[u].append((v, w))
            graph[v].append((u, w))

        visited = [False] * V
        pq = [(0, 0)]  # (weight, node)
        total = 0
        count = 0

        while pq and count < V:
            wt, node = heapq.heappop(pq)
            if visited[node]:
                continue

            visited[node] = True
            total += wt
            count += 1

            for nei, w in graph[node]:
                if not visited[nei]:
                    heapq.heappush(pq, (w, nei))

        return total

'''
 Given undirected weighted graph.
 with V nodes. and E edges.
 We need to find minimum spanning tree.
  What means min span tree:
    - all nodes must be connected, must can reach each others somehow.
    - the connected Edge weights must be minmium as possible.
  At the end need to return weight of mst.

 Take example 1:
 we re given 3 node and 3 edges,
 How we can get min span tree?
 Well if we take:
    0->1->2  = 5+3 = 8.
    0->2->1 = 1+3= 4. - MST.
    How you can see on each ans we remove on e edge, why?
Because to connect N nodes we need only N-1 edges.
 So, we try to use min wt edg for each node.
 and for each node only 1 edge. ignoring others.
 How we can do this?
  - start from any node.
  - take min wt on each process.
  - keep track MinHeap for node weights:
    wt int; from, to int;
  - keep visited array to avoid circle and not visit node second time.
  - keep mst list to keep track edges. if you need list.
  - keep sum variable as result.
    mst list:
        []{wt, node, parent}
  - when pq becomes empty return sum.

'''
