#include <bits/stdc++.h>

#define ll long long
#define endl "\n"
using namespace std;

void solve() {
    ll n;
    cin >> n;
    
    set<string> st;
    for (ll i = 0; i < n; i++) {
        string s;
        cin >> s;
        st.insert(s)        ;
    }
    
    for (auto name: st) {
        cout << name << endl;
    }
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    int tt; cin >> tt; while (tt--) solve();
}
