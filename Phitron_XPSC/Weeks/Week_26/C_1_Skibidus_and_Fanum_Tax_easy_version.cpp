#include <bits/stdc++.h>
using namespace std;
const int n = 2e5 + 5;
int a[n], b[2];
void solve() {
    int n, m;
    cin >> n >> m;
    for (int i = 1; i <= n; i++) cin >> a[i];
    for (int i = 1; i <= m; i++) cin >> b[i];

    a[1] = min(a[1], b[1] - a[1]);
    for (int i = 2; i <= n; i++) {
        a[i] = min(a[i], b[1] - a[i]);

        if (a[i] < a[i - 1]) {
            a[i] = b[1] - a[i];
        }
        
        if (a[i] < a[i - 1]) {
            cout << "NO" << '\n';
            return;
        }
    }

    cout << "YES" << '\n';
}
int main() {
    int t;
    cin >> t;
    while (t--)
        solve();
    return 0;
}