#include <bits/stdc++.h>
#define ll long long
#define endl "\n"
using namespace std;

void solve(){
    int n;
    cin >> n;

    if (n & 1) {
        cout << 0 << endl;
        return;
    }
    
    int cows = n / 4;
    cout << cows + 1 << endl;
}

int main(){
    ios_base::sync_with_stdio(false); cin.tie(NULL);
    int tc = 1;
    cin >> tc;
    while(tc--) solve();
}