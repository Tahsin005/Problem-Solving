n = input()
a = list(map(int, input().split()))
print('NO' if any(abs(a[i] - a[i + 1]) > 1 for i in range(int(n) - 1)) else 'YES')