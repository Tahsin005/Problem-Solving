#include<bits/stdc++.h>
#define ll long long
using namespace std;
#define all(x) x.begin(),x.end()

void solve () {
  ll n, k; cin >> n >> k;
	map<ll,ll> cnt;
	for (ll i = 0; i < n; i++) {
		ll x; cin >> x;
		cnt[x]++;
	}

	vector<ll> cur;
	for (auto i : cnt) {
		cur.push_back(i.second);
	}
	sort(all(cur));
	reverse(all(cur));
	ll answer = 0;
	while (cur.size() > 1 and cur.back() <= k)
	{
		k -= cur.back();
		cur.pop_back();
	}
	cout << cur.size() << "\n";
}

int main () {
  int tt; cin >> tt; while (tt--) solve();
  return 0;
}
