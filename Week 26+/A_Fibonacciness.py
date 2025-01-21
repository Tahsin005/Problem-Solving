n = int(input())
for i in range(n):
    a, b, d, e = map(int, input().split(' '))
    x = set([a + b, d - b, e - d])
    print(4 - len(x))