#include<bits/stdc++.h>
using namespace std;

int main () {
  int tt; cin >> tt; while (tt--) {
    int n; cin >> n;

    vector<int> a(n);
    for (int i = 0; i < n; i++) {
      cin >> a[i];
    }

    sort(a.begin(), a.end());

    int x = 0, idx = -1;

    for (int i = n - 2; i > -1; i--) {
      if (a[i] == a[i + 1]) {
        x = a[i];
        idx = i;
        break;
      }
    }

    if (idx == -1) {
      cout << -1 << '\n';
      continue;
    }

    a.erase(a.begin() + idx);
    a.erase(a.begin() + idx);

    bool found = false;

    for (int i = 0; i < n - 3; i++) {
      if (a[i] + 2 * x > a[i + 1]) {
        cout << a[i] << " " << x << " " << x << " " << a[i + 1] << '\n';
        found = true;
        break;
      }
    }

    if (!found) cout << -1 << '\n';
  }

  return 0;
}
