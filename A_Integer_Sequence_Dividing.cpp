#include <bits/stdc++.h>
using namespace std;
#define ll long long
const int M = 1e18+7;


void solve() {
    int n;
	cin >> n;
	n %= 4;
	if (n == 0 || n == 3) {
		cout << 0 << endl;
	} else {
		cout << 1 << endl;
	}
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt = 1; while (tt--) solve();
}
