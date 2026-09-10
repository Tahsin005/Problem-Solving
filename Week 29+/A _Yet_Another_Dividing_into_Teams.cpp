#include <bits/stdc++.h>
using namespace std;
#define ll long long

void solve() {
    int n;
    cin >> n;
    vector<int> a(n);
    for (int j = 0; j < n; ++j) {
        cin >> a[j];
    }
    bool ok = true;
    for (int p1 = 0; p1 < n; ++p1) {
        for (int p2 = 0; p2 < p1; ++p2) {
            ok &= abs(a[p1] - a[p2]) != 1;
		}
	}
	cout << 2 - ok << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt; cin >> tt; while (tt--) solve();
}
