#include <bits/stdc++.h>
using namespace std;
#define ll long long

void solve() {
    ll a, b, n, s;
    cin >> a >> b >> n >> s;

    ll use = min(a, s / n);

    ll remaining = s - use * n;

    if (remaining <= b) cout << "YES" << endl;
    else cout << "NO" << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt; cin >> tt; while (tt--) solve();
}
