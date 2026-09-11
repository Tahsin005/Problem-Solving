#include <bits/stdc++.h>
using namespace std;
#define ll long long


ll factorial(int n) {
    ll result = 1;
    for (int i = 2; i <= n; ++i) {
        result *= i;
    }
    return result;
}

ll combination(int n, int k) {
    if (k < 0 || k > n) return 0;
    return factorial(n) / (factorial(k) * factorial(n - k));
}

void solve() {
    int n;
    cin >> n;

    for (int i = 0; i < n; ++i) {
        int x;
        cin >> x;
    }

    ll ans = combination(10 - n, 2) * combination(4, 2);
    cout << ans << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt; cin >> tt; while (tt--) solve();
}
