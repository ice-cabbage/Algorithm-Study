sik_minus = input().split('-')
sik_plus = list(map(lambda x: sum(map(int,x.split('+'))),sik_minus))
print(sik_plus[0]-sum(sik_plus[1:]))