#include <bits/stdc++.h>
using namespace std;
 
void solve(){
    int n, x = 0, idx;
    cin >> n;
    vector<int> a(n), b(n);
    for(int i=0; i<n; i++){
        cin >> a[i];
        x ^= a[i];
    }
    for(int i=0; i<n; i++){
        cin >> b[i];
        x ^= b[i];
    }
    if(!x){
        cout << "Tie" << '\n';
        return;
    }
    for(int i=0; i<n; i++)
        if(a[i] ^ b[i])
            idx = i;
    cout << (idx & 1 ? "Mai" : "Ajisai") << '\n';
}
 
int main(){
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    int t;
    cin >> t;
    while(t--) solve();
}
