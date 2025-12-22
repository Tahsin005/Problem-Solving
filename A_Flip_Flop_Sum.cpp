#include <bits/stdc++.h>
using namespace std;
#define ll long long
const int M = 1e18+7;

void solve() {
    int n, sum = 0;
    cin >> n;
    
    vector<int> a(n + 1);
    for (int i = 1; i <= n; i++) {
        cin >> a[i];
        sum += a[i];
    }
    
    int ans = -1e9;
 
    for (int i = 1; i < n; i++) {
        if (a[i] == a[i + 1]) {
            if (a[i] == 1) ans = max(ans, sum - 4);
            else ans = max(ans, sum + 4);
        }
        else {
            ans = max(ans, sum);
        }
    }
    
    cout << ans << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt; cin >> tt;
    while (tt--) {
        solve();
    }
    
    return 0;
}
