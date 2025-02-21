#include<bits/stdc++.h>
using namespace std;

int main () {
  int n; 
  cin >> n;

  vector<int> a(n);
  for (int i = 0; i < n; i++) {
    cin >> a[i];
  }

  int sum = 0;
  for (int val : a) {
    sum += val;
  }

  cout << sum << endl;
  return 0;
}
