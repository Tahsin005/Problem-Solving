#include <bits/stdc++.h>
using namespace std;

long long f(long long x) {
    x += 1;
    while (x % 10 == 0) {
        x /= 10;
    }
    return x;
}

int main() {
    long long n;
    cin >> n;

    unordered_set<long long> a;

    while (!a.count(n)) {
        a.insert(n);
        n = f(n);
    }

    cout << a.size() << endl;
    return 0;
}
