#include <bits/stdc++.h>
using namespace std;
#define ll long long

bool isPossible(ll x, ll n, ll k) {
    ll smallestMultiple = ((n + k - 1) / k) * k;
    return smallestMultiple <= n * x;
}

void solve() {
    ll n, k;
    cin >> n >> k;

    ll low = 1, high = k, ans = k;

    while (low <= high) {
        ll mid = low + (high - low) / 2;
        if (isPossible(mid, n, k)) {
            ans = mid;
            high = mid - 1;
        } else {
            low = mid + 1;
        }
    }

    cout << ans << "\n";
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int t; cin >> t;
    while (t--) solve();
}
