#include <bits/stdc++.h>
using namespace std;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    int tt; cin >> tt; while (tt--) {
        int n;
        cin >> n;

        vector<int> a(n);

        for (int i = 0; i < n; i++) {
            cin >> a[i];
        }

        bool ok = false;

        for (int i = 0; i < n; i++) {
            int d = 2 * max(i, n - 1 - i);
            if (a[i] <= d) {
                ok = true;
                break;
            }
        }

        cout << (ok ? "NO" : "YES") << '\n';
    }

    return 0;
}
