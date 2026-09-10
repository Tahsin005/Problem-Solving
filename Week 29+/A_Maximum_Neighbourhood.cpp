#include <bits/stdc++.h>
using namespace std;
#define ll long long


void solve() {
    int n;
    cin >> n;
    int ans = 1;
    if (n == 2) ans = 9;
    else if (n == 3 || n == 4) ans = 4 * n * n - n - 4;
    else if (n > 4) ans = 5 * n * n - 5 * n - 5;
    cout << ans << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt; cin >> tt; while (tt--) solve();
}
