#include <bits/stdc++.h>
#define ll long long
#define endl "\n"
using namespace std;

void solve(){
    int n;
    cin >> n;
    int a;
    map<int, int> m;
    for (int i = 0; i < n; i++) cin >> a, m[a]++;
    int c = 0;
    vector<int> v;
    for (auto &it: m) {
        if (it.second == 1) {
            cout << -1 << endl;
            return;
        }
        for (int i = 0; i < it.second; i++) {
            v.push_back(c + (i + 1) % it.second + 1);
        }
        c += it.second;
    }
    for (int i = 0; i < n; i++) cout << v[i] << " ";
    cout << endl;
}

int main(){
    ios_base::sync_with_stdio(false); cin.tie(NULL);
    int tc=1;
    cin >> tc;
    while(tc--) solve();
}