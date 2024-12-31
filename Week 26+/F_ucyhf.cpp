#include <bits/stdc++.h>
using namespace std;
bool p(int n) {
    if (n < 2)
        return 0;
    for (int i = 2, m = sqrt(n); i <= m; i++)
        if (!(n % i))
            return 1;
    return 0;
}
main() {
    int n, a = 9, b;
    string s;
    cin >> n;
    while (n--) {
        do {
            while (p(++a))
                ;
            s = to_string(a), b = stoi(string(s.rbegin(), s.rend()));
        } while (a == b || p(b));
    }
    cout << a;
}