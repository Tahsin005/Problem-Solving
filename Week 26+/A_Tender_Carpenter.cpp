#include<bits/stdc++.h>
using namespace std;

int main () {
    int tt; cin >> tt; while (tt--) {
        int n;
        cin >> n;

        vector<int> a(n);
        for (int i = 0; i < n; i++) cin >> a[i];

        bool flag = false;

        for (int i = 0; i < n - 1; i++) {
            if (2 * min(a[i], a[i + 1]) > max(a[i], a[i + 1])) {
                flag = true;
                break;
            }
        }

        if (flag) cout << "YES" << '\n';
        else cout << "NO" << '\n';
    }
    return 0;
}