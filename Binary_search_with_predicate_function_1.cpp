#include <bits/stdc++.h>
using namespace std;
#define ll long long

const int N = 1e6+10;
int n;
ll m;
ll trees[N];

// https://www.spoj.com/problems/EKO/

bool isWoodSufficient(int height) {
    ll wood = 0;
    for (int i = 0; i < n; i++) {
        if (trees[i] >= height) {
            wood += (trees[i] - height);
        }
    }
    
    return wood >= m;
}

void solve() {
    cin >> n >> m;
    
    for (int i = 0; i < n; i++) {
        cin >> trees[i];
    }
    
    ll low = 0, high = 1e9, mid;
    // T T T T T F F F F F 
    while (high - low > 1) {
        mid = low + (high - low) / 2;
        if (isWoodSufficient(mid)) {
            low = mid;
        } else {
            high = mid - 1;
        }
    }
    
    if (isWoodSufficient(high)) {
        cout << high << endl;
    } else {
        cout << low << endl;
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    solve();
}
