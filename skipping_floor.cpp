#include <bits/stdc++.h>
using namespace std;
#define ll long long
const int M = 1e18+7;
 
void SUIII() {
    int n;
    cin >> n;
    
    ll a[n + 1];
    for (int i = 1; i <= n; i++) cin >> a[i];
    
    ll sum = 0;
    for (int i = 2; i <= n; i++) sum += abs(a[i] - a[i - 1]);
    
    ll ans = min(sum - abs(a[2] - a[1]), sum - abs(a[n] - a[n - 1]));
    for (int i = 2; i < n; i++) {
        ans = min(ans, sum - abs(a[i + 1] - a[i]) - abs(a[i] - a[i - 1]) + abs(a[i + 1] - a[i - 1]));
    }
    
    cout << ans << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt = 1; 
    cin >> tt;
    while (tt--) {
        SUIII();
    }
    
    return 0;
}
