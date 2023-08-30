MAX_HEIGHT = 256
MIN_HEIGHT = 0
N, M, B = map(int, input().split())
board = [ list(map(int, input().split())) for _ in range(N)]
height_cnt = [0 for _ in range(MAX_HEIGHT + 1)]
answer_cost = 100000001
answer_height = 0
sum_height = B 

for row in board:
    for val in row:
        height_cnt[val] += 1
        sum_height += val
     
min_h = min([ min(row) for row in board ])
max_h = sum_height // (M * N)

for h in range(min_h, max_h+1):
    cost = 0
    usable_blocks = B
    for i in range(MIN_HEIGHT, MAX_HEIGHT+1 ):
        if h < i:
            cost += height_cnt[i] * abs(h - i) * 2
            usable_blocks += height_cnt[i] * abs(h - i)
        else:
            cost += height_cnt[i] * abs(h - i)
            usable_blocks -= height_cnt[i] * abs(h - i)
            
    if cost <= answer_cost and usable_blocks >= 0:
        answer_cost = cost
        answer_height = h

print(answer_cost, answer_height)