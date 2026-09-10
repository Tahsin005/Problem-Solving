#include <bits/stdc++.h>
using namespace std;
#define ll long long


void solve() {
    int n, a;
    cin >> n >> a;
    vector <int> v(n);
    int l = 0, r = 0;
    for (int i = 0; i < n; ++i) {
        cin >> v[i];
        if (a > v[i]) ++l;
        if (a < v[i]) ++r;
    }
        
    cout << (l > r ? a - 1 : a + 1) << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt; cin >> tt; while (tt--) solve();
}
