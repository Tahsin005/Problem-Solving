#include <bits/stdc++.h>
#define ll long long
#define endl "\n"
using namespace std;

void solve(){
    ll q, y, x;
    cin >> q;
    
    set<ll> s;
    while (q--) {
        cin >> y >> x;
        if (y == 1) s.insert(x);
        else if (y == 2) {
            auto it = s.find(x);
            if (it != s.end()) {
                s.erase(it);
            }
        } else {
            auto it = s.find(x);
            if (it != s.end()) cout << "Yes" << endl;
            else cout << "No" << endl;
        }
    }
}

int main(){
    ios_base::sync_with_stdio(false); cin.tie(NULL);
    solve();
}
