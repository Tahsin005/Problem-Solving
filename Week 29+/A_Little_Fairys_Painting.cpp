#include <bits/stdc++.h>
using namespace std;
#define ll long long
const int M = 1e18+7;

void solve() {
    int n;
    cin >> n;
    vector<int>a (n);
    set<int> s;
    for (int i = 0; i < n; i++) {
        cin >> a[i];
        s.insert(a[i]);
    }
    sort(a.begin(), a.end());
    auto it = lower_bound(a.begin(), a.end(), s.size());
    cout<< *it <<endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt = 1; cin >> tt; while (tt--) solve();
}
