import sys
from collections import deque
sys.setrecursionlimit(1000000)
input = sys.stdin.readline

def dfs(graph, x, y):
    if x <= -1 or x >= N or y <= -1 or y >= M:
        return False
    
    if graph[x][y] == 1:
        graph[x][y] = 0
        
        dfs(graph, x - 1, y)
        dfs(graph, x + 1, y)
        dfs(graph, x, y - 1)
        dfs(graph, x, y + 1)
        return True
    return False

T = int(input())

for _ in range(T):
    M, N, K = map(int, input().split())
    g = [[0] * M for _ in range(N)]
    for _ in range(K):
        i, j = map(int, input().split())
        g[j][i] = 1
        
    worm = 0
    for i in range(N):
        for j in range(M):
            if dfs(g, i, j) == True:
                worm += 1
                
    print(worm)