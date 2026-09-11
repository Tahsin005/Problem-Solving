#include <bits/stdc++.h>
using namespace std;

int main() {
    string s;
    cin >> s;

    size_t a = s.find('Q');
    if (a == string::npos) {
        cout << "Yes" << '\n';
        return 0;
    }

    a /= 2;
    int cnt = count(s.begin(), s.end(), 'Q');
    int sqc = static_cast<int>(sqrt(cnt));
    int w = (s.size() - cnt) / (sqc + 1);

    string b = s.substr(a, w + sqc);
    string ans;
    for (char ww : b) {
        if (ww == 'H') {
            ans += 'H';
        } else {
            ans += b;
        }
    }

    if (s == ans) {
        cout << "Yes" << '\n';
    } else {
        cout << "No" << '\n';
    }

    return 0;
}
