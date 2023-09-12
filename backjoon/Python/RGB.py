n = int(input())

r, g, b = map(int, input().split())
red = [r]
green = [g]
blue = [b]

for i in range(1, n):
    r, g, b = map(int, input().split())
    red.append(r + min(green[i-1], blue[i-1]))
    green.append(g + min(red[i-1], blue[i-1]))
    blue.append(b + min(red[i-1], green[i-1]))

print(min(red[-1], blue[-1], green[-1]))