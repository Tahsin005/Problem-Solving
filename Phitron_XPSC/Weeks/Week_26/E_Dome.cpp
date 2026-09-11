#include <bits/stdc++.h>
using namespace std;
int main() {
    double x;
    cin >> x;
    for (int i = 1; i <= 10; i++)
        for (int j = 1; j <= 10; j++)
            if (fabs(i * j / sqrt(i * i + 4 * j * j) - x) <= 1e-5)
                return printf("%d %d\n", i, j), 0;
}