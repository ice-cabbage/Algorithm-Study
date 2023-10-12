import sys

input_sys = sys.stdin.readline
n, k = map(int, input().split())
nums = list(map(int, input().split()))
sums = sum(nums[:k])
li = [sums]

for i in range(len(nums)-k):
    sums = sums - nums[i] + nums[i+k]
    li.append(sums)

print(max(li))