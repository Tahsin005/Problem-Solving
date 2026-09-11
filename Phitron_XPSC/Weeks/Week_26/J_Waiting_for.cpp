#include<bits/stdc++.h>
#define ll long long
using namespace std;

int main() {
    int n, acc = 0, a;
    char ch;

    cin >> n;

    while (n--) {
        cin >> ch >> a;

        if (ch == 'P') {
            acc += a;
        } else if (a > acc) {
            acc = 0;
            cout << "YES" << endl;
        } else {
            acc -= a;
            cout << "NO" << endl;
        }
    }

    return 0;
}