#include <bits/stdc++.h>

#define ll long long
#define endl "\n"
using namespace std;

void solve() {
    ll n, m;
    cin >> n >> m;
    
    map<ll, ll> mp;
    for (ll i = 0; i < n; i++) {
        ll x;
        cin >> x;
        mp[x]++;
    }
    
    for (ll i = 0; i < m; i++) {
        ll x;
        cin >> x;
        
        if (mp.find(x) != mp.end()) {
            cout << "YES" << endl;
        } else {
            cout << "NO" << endl;
        }
        
        mp[x]++;
    }
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    int tt; cin >> tt; while (tt--) solve();
}
