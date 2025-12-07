#include <bits/stdc++.h>
using namespace std;
#define ll long long
const int M = 1e9+7;

int sum(int a) {
    int result = 0;
    while (a > 0) {
        result += a % 10;
        a /= 10;
    }
    return result;
}

void solve() {
    int a;
    cin >> a;
    while (sum(a) % 4 != 0) {
        a++;
    }
    cout << a << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    solve();
}
