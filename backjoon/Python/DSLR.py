from collections import deque
import sys
input = sys.stdin.readline

T = int(input())
for _ in range(T):
    A, B = map(int, input().split())

    visit = [False]*10000
    queue = deque([(A, "")])
    visit[A] = True
    while queue:
        n, x = queue.popleft() 

        if n == B:
            print(x)
            break

        num = (n*2) % 10000
        if not visit[num]:
            visit[num] = True
            queue.append((num, x + 'D'))

        num = (n-1) if n != 0 else 9999 
        if not visit[num]:
            visit[num] = True 
            queue.append((num, x + 'S'))

        num = (n%1000)*10 + n//1000
        if not visit[num]:
            visit[num] = True 
            queue.append((num, x + 'L'))

        num = n//10 + (n%10)*1000 
        if not visit[num]:
            visit[num] = True
            queue.append((num, x + 'R'))