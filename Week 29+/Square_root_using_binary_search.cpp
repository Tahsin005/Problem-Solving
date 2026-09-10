#include <bits/stdc++.h>
using namespace std;
#define ll long long

double eps = 1e-6;

void solve() {
    double x; cin >> x;
   
    double low = 1, high = x, mid;
    while (high - low > eps) {
        mid = (high + low) / 2;
        if (mid * mid < x) {
            low = mid;
        } else {
            high = mid;
        }
    }
   
   cout << low << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    solve();
}
