n = int(input())
hansu = 0

for i in range(1, 1 + 1):
    if i <= 99:
        hansu += 1
        
    else:
        num = list(map(int, str(i)))
        if num[2] - num[1] == num[1] - num[0]:
            hansu += 1
            
print(hansu)