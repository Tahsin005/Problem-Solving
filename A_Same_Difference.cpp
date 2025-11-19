#include <bits/stdc++.h>
#define ll long long
#define endl "\n"
using namespace std;

void solve(){
    int n, cnt = 0;
    cin >> n;
    string s;
    cin >> s;
    
    char char_to_compare = s[s.length() - 1];
    for (char ch: s) {
        cnt += (ch != char_to_compare);
    }
    
    cout << cnt << endl;
}

int main(){
    ios_base::sync_with_stdio(false); cin.tie(NULL);
    int tt; cin >> tt; while (tt--) solve();
}
