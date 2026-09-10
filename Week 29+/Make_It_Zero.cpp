#include <bits/stdc++.h>

#define ll long long
#define endl "\n"
using namespace std;

void solve() {
    int n; cin >> n;
	vector<int> a(n);
	for (int i = 0; i < n; i++) cin >> a[i];
	if (n & 1) {
		cout << "4" << endl;
		cout << "1 " << n - 1 << endl;
		cout << "1 " << n - 1 << endl;
		cout << n - 1 << ' ' << n << endl;
		cout << n - 1 << ' ' << n << endl;
	} else {
		cout << "2" << endl;
		cout << "1 " << n << endl;
		cout << "1 " << n << endl;
	}
	return;
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    int tt; cin >> tt; while (tt--) solve();
}
