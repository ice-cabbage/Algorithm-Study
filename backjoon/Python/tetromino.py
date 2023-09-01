n, m = map(int, input().split(' '))

maps = [list(map(int, input().split(' '))) for _ in range(n)]
check = [[True]*m for _ in range(n)]

dx = [-1, 0, 1, 0]
dy = [0, -1, 0, 1]
max_result = -1

def go(x, y, now, cnt):
    global max_result
    
    if cnt == 4:
        if now > max_result:
            max_result = now
        return

    for i in range(4):
        new_x = x + dx[i]
        new_y = y + dy[i]

        if 0 <= new_x < n and 0 <= new_y < m:
            if check[new_x][new_y]:
                check[new_x][new_y] = False
                go(new_x, new_y, now+maps[new_x][new_y], cnt+1)
                check[new_x][new_y] = True

def nogo(x, y):
    center = maps[x][y]
    
    wings = 4

    min_result = 100000

    for i in range(4):
        new_x = x + dx[i]
        new_y = y + dy[i]
        
        if wings == 2:
            return 0
        
        if new_x < 0 or new_x >= n or new_y < 0 or new_y >= m:
            wings -= 1
            continue
        
        center += maps[new_x][new_y]
        if maps[new_x][new_y] < min_result:
            min_result = maps[new_x][new_y]
    
    if wings == 4:
        center -= min_result

    return center


for i in range(n):
    for j in range(m):
        check[i][j] = False
        go(i, j, maps[i][j], 1)
        check[i][j] = True

        temp = nogo(i, j)
        if temp > max_result:
            max_result = temp

print(max_result)