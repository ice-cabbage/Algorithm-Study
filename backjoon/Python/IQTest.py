import sys

def solve(a, b):
    for i in range(N-1):
        
        if nums[i] * a + b != nums[i+1]:
            return False 

    return True


if __name__ == '__main__':
    N = int(input())
    nums = list(map(int, input().split()))

    if N == 1:
        print('A')
        sys.exit(0)

    if N == 2:
        if nums[0] == nums[1]: 
            print(nums[0])
        else: 
            print('A')
        sys.exit(0)

    if nums[0] == nums[1] == nums[2]:
        if nums.count(nums[0]) == N:
            print(nums[0])
        else: 
            print('B')
        sys.exit(0)

    a, b = 0, 0
    
    for i in range(N-2): 
        if nums[i+1] - nums[i] == 0:
            continue
        else: 
            a = (nums[i+2] - nums[i+1]) // (nums[i+1] - nums[i])
            b = nums[i+1] - (nums[i] * a)
            break
  
    if solve(a, b):
        print(nums[N-1] * a + b)
    else:
        print('B')