from collections import deque
N = int(input())
student = list(map(int,input().split()))
stack = []
idx = 0
end = 1
while idx < N:
	if stack and stack[-1] == end:
		end += 1
		stack.pop()
	else:
		stack.append(student[idx])
		idx += 1
 
if stack:
	while stack:
		if stack.pop() == end:
			end += 1
		else:
			print("Sad")
			break
 
if end == N+1:
	print("Nice")