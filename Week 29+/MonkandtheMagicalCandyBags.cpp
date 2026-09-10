#include <bits/stdc++.h>
#define ll long long
#define endl "\n"
using namespace std;

// https://www.hackerearth.com/practice/data-structures/trees/heapspriority-queues/practice-problems/algorithm/monk-and-the-magical-candy-bags/

void solve(){
    ll n, k;
    cin >> n >> k;
    
    multiset<ll> s;
    for (ll i = 0; i < n; i++) {
        ll x;
        cin >> x;
        
        s.insert(x);
    }
    
    ll cnt = 0;
    while (k--) {
        if (s.size() > 0) {
            auto it = prev(s.end());
            cnt += (*it);
            
            ll new_value = (*it) / 2;
            s.erase(it);
            s.insert(new_value);
        }
    }
    
    cout << cnt << endl;
}

int main(){
    ios_base::sync_with_stdio(false); cin.tie(NULL);
    int tt; cin >> tt; while (tt--) solve();
}
