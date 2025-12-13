#include <bits/stdc++.h>
using namespace std;
#define ll long long
const int M = 1e18+7;

void solve() {
    int n, k;
    cin >> n >> k;
    
    string s;
    cin >> s;
    
    int ans = 0;
    for (int i = 0; i < n; i++) {
        if (s[i] == '0') {
            int ok = 1;
            for (int j = max(0, (int)i - k); j < i; j++) {
                if (s[j] == '1') {
                    ok = 0;
                    break;
                }
            }
            ans += ok;
        }
    }
    
    cout << ans << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt = 1; cin >> tt; while (tt--) solve();
}
