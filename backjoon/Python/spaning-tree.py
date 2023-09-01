import sys
from math import log2
from collections import deque

def find(node):
    if parent[node] != node:
        parent[node] = find(parent[node])
    return parent[node]

def union(a, b):
    parent_a = find(a)
    parent_b = find(b)

    if parent_a > parent_b:
        parent[parent_a] = parent_b
    else:
        parent[parent_b] = parent_a

def bfs():
    depth[0] = 0
    q = deque([0])

    while q:
        node = q.popleft()

        for next_node, cost in graph[node]:
            if depth[next_node] != -1:
                continue

            depth[next_node] = depth[node] + 1
            parent[next_node][0] = node
            two_weight[next_node][0] = [cost, -1]
            q.append(next_node)

def longest(arr1, arr2):
    temp = list(set(arr1 + arr2))
    temp.sort(reverse=True)

    while len(temp) < 2:
        temp.append(-1)

    temp = temp[:2]
    return temp

def set_parent():
    bfs()

    for log in range(1, k):
        for node in range(1, v):
            next_node = parent[node][log - 1]
            parent[node][log] = parent[next_node][log - 1]

            weight1, weight2 = two_weight[node][log - 1], two_weight[next_node][log - 1]
            two_weight[node][log] = longest(weight1, weight2)

def lca(a, b):
    ret = [-1, -1]
    if depth[a] > depth[b]:
        a, b = b, a

    for log in range(k - 1, -1, -1):
        if depth[b] - depth[a] >= (1 << log):
            ret = longest(ret, two_weight[b][log])
            b = parent[b][log]

    if a == b:
        return ret

    for log in range(k - 1, -1, -1):
        if parent[b][log] != parent[a][log]:
            ret = longest(ret, two_weight[a][log])
            ret = longest(ret, two_weight[b][log])
            b = parent[b][log]
            a = parent[a][log]

    ret = longest(ret, two_weight[a][0])
    ret = longest(ret, two_weight[b][0])

    return ret


v, e = map(int, sys.stdin.readline().split())
edge, graph, used = [], [[] for i in range(v)], [0] * e
parent = [i for i in range(v)]

for i in range(e):
    a, b, c = map(int, sys.stdin.readline().split())
    a -= 1
    b -= 1
    edge.append((c, a, b))

edge.sort()
mst, cnt = 0, 0

for i in range(e):
    c, a, b = edge[i]

    if find(a) != find(b):
        union(a, b)
        graph[a].append((b, c))
        graph[b].append((a, c))

        used[i] = 1
        mst += c
        cnt += 1

if cnt != v - 1:
    print(-1)
    exit(0)

ans, k = float("inf"), int(log2(v)) + 1
depth, two_weight = [-1] * v, [[[-1, -1] for _ in range(k)] for _ in range(v)]
parent = [[-1] * k for i in range(v)]

set_parent()
for i in range(e):
    if used[i]:
        continue

    w, u, v = edge[i]
    weight = lca(u, v)

    if weight[0] != w:
        ans = min(ans, mst - weight[0] + w)
    elif weight[1] != w and weight[1] != -1:
        ans = min(ans, mst - weight[1] + w)

if ans == float("inf"):
    print(-1)
    exit(0)

print(ans)