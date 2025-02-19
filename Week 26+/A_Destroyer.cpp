#include<bits/stdc++.h>
using namespace std;
#define el '\n'
void solve() {
	long long n; cin >> n;

	map<long long, long long> mp;
	for (long long i = 0; i < n; i++) {
		long long x; cin >> x; mp[x]++;
	}

	bool okay = true;

	for (int i = 0; i < n; i++) {
		if (mp[i] >= mp[i + 1]) continue;
		else {
			okay = false;
			break;
		}
	} 

	if (okay) return void(cout << "YES" << el);
	else return void(cout << "NO" << el);
}
int main () {
	long long tt; cin >> tt; while(tt--) {
		solve();
	}
}
