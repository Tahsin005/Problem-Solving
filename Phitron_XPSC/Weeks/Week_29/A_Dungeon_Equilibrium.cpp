#include <bits/stdc++.h>
using namespace std;
#define ll long long


void solve() {
    int n; cin >> n;

    vector<int> freq(110, 0);
    for (int i = 0; i < n; i++) {
        int x; cin >> x;
        freq[x]++;
    }
    
    int cnt = 0;
    for (int i = 0; i < 110; i++) {
        if (freq[i] > 0) {
            if (freq[i] != i) {
                if (freq[i] > i) {  
                    cnt += abs(freq[i] - i);
                } else {
                    cnt += freq[i];
                }
            }
        }
    }
    
    cout << cnt << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt; cin >> tt; while (tt--) solve();
}
