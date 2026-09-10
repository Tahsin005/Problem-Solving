#include <bits/stdc++.h>
using namespace std;
#define ll long long
const int M = 1e9+7;

int binExpIter(int a, int b) {
    int ans = 1;
    while (b) {
        if (b & 1) {
            ans = (ans * 1LL * a) % M;
        }
        a = (a * 1LL * a) % M;
        b >>= 1;
    }
    
    return ans;
}

void solve() {
    int a = 2, b = 13;
    cout << binExpIter(a, b) << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt; cin >> tt; while (tt--) solve();
}
