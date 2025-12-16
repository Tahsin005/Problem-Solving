#include <bits/stdc++.h>
#define ll long long
#define endl "\n"
using namespace std;

void solve(){
    int n;
	cin >> n;
	for (int i = 0; i < n; ++i) {
		int x;
		cin >> x;
		cout << x - !(x & 1) << " ";
	}
}

int main(){
    ios_base::sync_with_stdio(false); cin.tie(NULL);
    solve();
}