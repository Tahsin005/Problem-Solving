#include<bits/stdc++.h>
using namespace std;

int main () {
  int tt; cin >> tt; while (tt--) {
    int n, x, y = 0;
    cin >> n >> x;

    for (int i = 1; i <= 3; i++) {
      vector<int> v(n);
      for (int j = 0; j < n; j++) cin >> v[j];

      for (int j = 0; j < n; j++) {
        if ((x | v[j]) == x) y |= v[j];
        else break;
      }
    }

    if (x == y) cout << "YES\n";
    else cout << "NO\n";
  }
}
