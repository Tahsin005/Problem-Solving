#include <bits/stdc++.h>
#define ll long long
#define endl "\n"
using namespace std;

int a[10005];
void solve(){
    int n;
    cin >> n;
    for (int i = 1; i <= n; i++)
        cin >> a[i];
    int p = 0;
    for (int i = 2; i <= n - 1; i++){  
        if (a[i] > a[i - 1] and a[i] > a[i + 1])
            p = i;
    }
    if (p)
        cout << "YES" << endl
             << p - 1 << " " << p << " " << p + 1 << endl;
    else
        cout << "NO" << endl;
}

int main(){
    ios_base::sync_with_stdio(false); cin.tie(NULL);
    int tc=1;
    cin >> tc;
    while(tc--) solve();
}