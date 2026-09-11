#include <bits/stdc++.h>
using namespace std;

void solve(){
    long long n, m;
    cin >> n >> m;
    pair<long long, long long> a[m + 2];
    for (long long i = 1; i <= m; i++) {
        cin >> a[i].second >> a[i].first;
    }
    sort(a + 1, a + m + 1);
    long long k = m, s = 0;
    while (k >= 1) {
        if (n > a[k].second) {
            s += a[k].second * a[k].first;
            n -= a[k].second;
        } else{
            s += n * a[k].first;
            break;
        }
        k--;
    }
    cout << s;
}

int main(){
    ios_base::sync_with_stdio(false); cin.tie(NULL);
    solve();
}