#include <bits/stdc++.h>
using namespace std;

int main() {
    int t, r, l, R, L;
    cin >> t;
    while (t--) {
        cin >> l >> r >> L >> R;
        cout << max(min(r, R) - max(l, L) + (R != r) + (l != L), 1) << endl;
    }
}