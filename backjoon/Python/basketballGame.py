N = int(input())
lst = [0] * 26

for _ in range(N):
    name = input()
    lst[ord(name[0]) - 97] += 1

if max(lst) >= 5:
    for i in range(len(lst)):
        if lst[i] >= 5:
            print(chr(i + 97), end = '')
else:
    print("PREDAJA")