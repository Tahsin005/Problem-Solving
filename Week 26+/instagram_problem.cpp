#include<bits/stdc++.h>
using namespace std;

int main () {
  string s, tmp = "", ans = "";
  cin >> s;
  char lastInt = '#';
  bool start = false;
  for (int i = 0; i < s.length(); i++) {
    if (s[i] >= '0' and s[i] <= '9') {
      lastInt = s[i];
    } else if (s[i] == '[') {
      start = true;
    } else if (start == true and (s[i] >= 'a' and s[i] <= 'z')) {
      tmp += s[i];
    } else if (tmp != "" and s[i] == ']' and lastInt != '#') {
      int x = lastInt - '0';

      while (x--) {
        ans += tmp;
      }

      tmp = "";
      lastInt = '#';
      start = false;
    }
  }

  cout << ans << endl;
  return 0;
}
