#include <bits/stdc++.h>
#define ll long long
#define endl "\n"
using namespace std;

void solve(){
    ll x0, n, d;
    cin >> x0 >> n;
    
    ll r = n % 4;
    
    if (r == 0) d = 0;
    else if (r == 1) d = n;
    else if (r == 2) d = -1;
    else if (r == 3) d = -n - 1;
    
    if (x0 % 2 == 0) cout << x0 - d << endl;
    else cout << x0 + d << endl;
}

int main(){
    ios_base::sync_with_stdio(false); cin.tie(NULL);
    int tt; cin >> tt; while (tt--) solve();
}
