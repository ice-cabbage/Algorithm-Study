from collections import deque

N, M = map(int, input().split())
graph =[]
for i in range(N):
    graph.append(list(map(int, input())))

visited = [[[0] * 2 for _ in range(M)] for _ in range(N)]
visited[0][0][1] = 1

def bfs():
    queue = deque([[0, 0, 1]])
    dx = [-1, 1, 0, 0]
    dy = [0, 0, -1, 1]

    while queue:
        y, x, w = queue.popleft()
        if y == N-1 and x == M-1:
            return visited[y][x][w]

        for i in range(4):
            nx = x + dx[i]
            ny = y + dy[i]

            if 0<=nx<M and 0<=ny<N:
                if graph[ny][nx] == 0 and visited[ny][nx][w] == 0:
                    visited[ny][nx][w] = visited[y][x][w] + 1
                    queue.append([ny, nx, w])

                if graph[ny][nx] == 1 and w == 1:
                    visited[ny][nx][0] = visited[y][x][1] + 1
                    queue.append([ny, nx, 0])
    return -1
print(bfs())