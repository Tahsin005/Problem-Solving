#include <bits/stdc++.h>
using namespace std;
#define ll long long

void solve() {
    int n, k; 
    cin >> n >> k;
    for (int i = 0; i < n - k; i++) cout << 0;
    for (int i = 0; i < k; i++) cout << 1;
    cout << "\n";
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt; cin >> tt; while (tt--) solve();
}
