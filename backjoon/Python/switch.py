import sys
import math
input = sys.stdin.readline


def get_segment_tree_length():
    if N & (N-1) == 0:
        return 2*N
    else:
        return pow(2, math.ceil(math.log(N, 2)) + 1)


def update_segment_tree(index, start, end, left, right):
    propagate_segment_tree(index, start, end)

    # Out of range
    if right < start or end < left:
        return

    # In range
    if left <= start and end <= right:
        seg_tree[index] = (end - start + 1) - seg_tree[index]

        if start != end:
            lazy[index*2] ^= 1
            lazy[index*2+1] ^= 1

        return

    mid = (start + end)//2
    update_segment_tree(index*2, start, mid, left, right)
    update_segment_tree(index*2+1, mid+1, end, left, right)
    seg_tree[index] = seg_tree[index*2] + seg_tree[index*2+1]


def query_segment_tree(index, start, end, left, right):
    propagate_segment_tree(index, start, end)

    # Out of range
    if right < start or end < left:
        return 0

    # In range
    if left <= start and end <= right:
        return seg_tree[index]

    mid = (start + end)//2
    return query_segment_tree(index*2, start, mid, left, right) + query_segment_tree(index*2+1, mid+1, end, left, right)


def propagate_segment_tree(index, start, end):
    if lazy[index] != 0:
        seg_tree[index] = (end - start + 1) - seg_tree[index]

        if start != end:
            lazy[index*2] ^= 1
            lazy[index*2+1] ^= 1

        lazy[index] = 0


if __name__ == '__main__':
    N, M = map(int, input().split())

    length = get_segment_tree_length()
    seg_tree = [0]*length
    lazy = [0]*length

    for _ in range(M):
        o, s, t = map(int, input().split())
        if o == 0:
            update_segment_tree(1, 1, N, s, t)
        else:
            print(query_segment_tree(1, 1, N, s, t))
