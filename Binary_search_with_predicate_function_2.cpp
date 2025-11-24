#include <bits/stdc++.h>
using namespace std;
#define ll long long

const int N = 1e5+10;
int n;
ll c;
ll stalls[N];

// https://www.spoj.com/problems/AGGRCOW/

bool canPlace(int distance) {
    ll curr = stalls[0];
    ll placed = 1;
    
    for (int i = 1; i < n; i++) {
        if (stalls[i] - curr >= distance) {
            placed++;
            curr = stalls[i];
            if (placed == c) return true;
        }
    }
    
    return false;
}

void solve() {
    cin >> n >> c;
    
    for (int i = 0; i < n; i++) {
        cin >> stalls[i];
    }
    
    sort(stalls, stalls + n);
    
    ll low = 1, high = 1e9, ans = 0;
    
    while (low <= high) {
        ll mid = low + (high - low) / 2;
        
        if (canPlace(mid)) {
            ans = mid;       
            low = mid + 1;
        } else {
            high = mid - 1;
        }
    }
    
    cout << ans << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt; cin >> tt; while (tt--) solve();
}
