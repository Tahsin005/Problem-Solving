#include <bits/stdc++.h>
using namespace std;
#define ll long long
const int M = 1e18+7;


void solve() {
    int n;
    cin >> n;
    int cnt = 0, sum = 0;
    
    for (int i = 1; i * i <= n; i++) {
        if (n % i == 0) {
            cout << i << " " << n / i << endl;
            cnt += 1;
            sum += i;
            if (n / i != i) {
                sum += n / i;
                cnt += 1;
            }
        }
    }
    
    cout << "Sum and count of divisors is: " << sum << " " << cnt << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt = 1; while (tt--) solve();
}
