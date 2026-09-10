#include <bits/stdc++.h>
#define ll long long
#define endl "\n"
using namespace std;

void solve(){
	int n, k;
	string t;
	cin >> n >> k >> t;
	
	int cnt = 1;
	int pos = 1;
	string ans = t;
	while (cnt < k) {
		if (pos >= ans.length()) {
			ans += t;
			++cnt;
		} else {
			bool ok = true;
			int len = 0;
			for (int i = 0; i < t.length(); ++i) {
				if (pos + i >= ans.length()) break;
				++len;
				if (ans[pos + i] != t[i]) ok = false;
			}
			if (ok) {
				ans += t.substr(len);
				++cnt;
			}
		}
		++pos;
	}
	
	cout << ans << endl;
}

int main(){
    ios_base::sync_with_stdio(false); cin.tie(NULL);
    solve();
}