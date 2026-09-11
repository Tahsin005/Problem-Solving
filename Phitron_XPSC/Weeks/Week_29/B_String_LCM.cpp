#include <bits/stdc++.h>
using namespace std;
#define ll long long

string mul(string s, int k) {
    string res = "";
    while (k--) {
        res += s;
    }
    return res;
}

void solve() {
    string s, t;
    cin >> s >> t;
    
    int n = s.length(), m = t.length();
    int g = __gcd(n, m);
    if (mul(s, m / g) == mul(t, n / g)) {
        cout << mul(s, m / g) << endl;
    } else {
        cout << -1 << endl;
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt; cin >> tt; while (tt--) solve();
}
