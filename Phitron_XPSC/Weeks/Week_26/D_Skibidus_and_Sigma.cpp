#include<bits/stdc++.h>
using namespace std;
using ll = long long;

void solve() {
  ll n, m; cin >> n >> m;

  vector<vector<ll>> ve(n, vector<ll>(m));

  for (vector<ll> &ve2 : ve) {
    for (ll &val : ve2) cin >> val;
  }

  vector<ll> vsum(n, 0);

  for (ll i = 0; i < n; i++) {
    for (ll j : ve[i]) vsum[i] += j;
  }
  
  vector<ll> th(n);
  iota(th.begin(), th.end(), 0);
  sort(th.begin(), th.end(), [&](ll a, ll b) {
      return vsum[a] > vsum[b];
  });

  ll ans = 0;
  for (ll i = 0; i < n; i++) {
    ans += vsum[th[i]] * (n - 1 - i) * m;
  }

  for (vector<ll> ve2 : ve) {
    for (ll i = 0; i < m; i++) {
      ans += ve2[i] * (m - i);
    }
  }

  cout << ans << '\n';
}
int main () {
  cin.tie(nullptr) -> sync_with_stdio(false);
  int tt; cin >> tt; while (tt--) {
    solve();
  }
  return 0;
}
