import sys

for i in range(3):
    a, b, c, d, e, f = map(int, input().split())
    start = a * 3600 + b * 60 +c
    end = d * 3600 + e * 60 + f
    gap = end - start
    
    h = gap // 3600
    m = (gap % 3600) // 60
    s = (gap % 3600) % 60
    
    print(h, m, s)