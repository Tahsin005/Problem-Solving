#include <bits/stdc++.h>
using namespace std;
#define ll long long


void solve() {
    int n, y, r;
    cin >> n >> y >> r;
    
    int cnt = r;
    cnt += (y / 2);
    cout << min(n, cnt) << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt; cin >> tt; while (tt--) solve();
}
