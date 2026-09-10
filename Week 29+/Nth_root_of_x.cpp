#include <bits/stdc++.h>
using namespace std;
#define ll long long

double eps = 1e-6;

double mulitply(double mid, int n) {
    double ans = 1;
    for (int i = 0; i < n; i++) {
        ans *= mid;
    }
    return ans;
}

void nth_root() {
    double x; cin >> x;
    int n; cin >> n;
   
    double low = 1, high = x, mid;
    while (high - low > eps) {
        mid = (high + low) / 2;
        if (mulitply(mid, n) < x) {
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

    nth_root();
}
