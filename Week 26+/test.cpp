#include<bits/stdc++.h>
using namespace std;

int main () {
  int n;
  cin >> n;

  vector<int> v(n);
  for (int i = 0; i < n; i++) cin >> v[i];
  int target;
  cin >> target;
  sort(v.begin(), v.end());

  bool found = false;

  int l = 0, r = n - 1;
  while (l < r) {
    int mid = (l + r) / 2;
    if (v[mid] == target) {
      found = true;
      break;
    } else if (v[mid] < target) {
      l = mid + 1;
    } else {
      r = mid - 1;
    }
  }

  cout << (found ? "YES" : "NO") << '\n';
  return 0;
}
