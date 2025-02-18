#include<bits/stdc++.h>
using namespace std;
void solve() {
	int n; cin >> n;

	map<int, int> mp;
	for (int i = 0; i < n; i++) {
		int x; cin >> x;
		mp[x]++;
	}
	bool ok = true;

	for (int i = 0; i < 99; i++) {
		if (mp[i] >= mp[i + 1]) continue;
		else {
			ok = false;
			break;
		}
	}

	if (ok) cout << "YES" << '\n';
	else cout << "NO" << '\n';
}
int main () {
	int tt; cin >> tt; while (tt--) {
		solve();
	}
}
