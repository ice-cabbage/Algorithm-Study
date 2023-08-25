import sys
input = sys.stdin.readline

N = int(input())
li = list(map(int, input().split()))
left, right = 0, N - 1
ans = [li[left], li[right]]

while right-1 != left:
    temp = li[left] + li[right]
    if temp < 0:
        left += 1
    else:
        right -= 1
    if abs(sum(ans)) >= abs(li[left] + li[right]):
        ans = [li[left], li[right]]
        
print(ans[0], ans[1])