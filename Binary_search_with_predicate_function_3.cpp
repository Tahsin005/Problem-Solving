#include <bits/stdc++.h>
using namespace std;
#define ll long long

const double PI = 3.141592653589793238;

double volumes[10005];

// https://www.spoj.com/problems/PIE/

bool can(double mid, int N, int people) {
    int cnt = 0;
    for (int i = 0; i < N; i++) {
        cnt += (int)(volumes[i] / mid);
        if (cnt >= people) return true;
    }
    return false;
}

void solve() {
    int N, F;
    cin >> N >> F;

    vector<int> r(N);
    for (int i = 0; i < N; i++) {
        cin >> r[i];
        volumes[i] = PI * r[i] * r[i];
    }
    
    int people = F + 1;

    double low = 0, high = *max_element(volumes, volumes + N);

    for (int i = 0; i < 60; i++) {
        double mid = (low + high) / 2.0;

        if (can(mid, N, people))
            low = mid;
        else
            high = mid;
    }

    cout << fixed << setprecision(4) << low << endl;;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt; cin >> tt; while (tt--) solve();
}
