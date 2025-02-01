#include<bits/stdc++.h>
#define ll long long
using namespace std;
/*



*/

int main() {
    int n;
    cin >> n;
    int a[n], b[n - 1], c[n - 2];
    for (int i = 0; i < n; i++) {
        cin >> a[i];
    }
    sort(a, a + n);
    for (int i = 0; i < n - 1; i++) {
        cin >> b[i];
    }
    sort(b, b + n - 1);
    int ok = 1;
    for (int i = 0; i < n - 1; i++) {
        if (a[i] != b[i]) {
            cout << a[i] << endl;
            ok = 0;
            break;
        }
    }
    if (ok == 1) {
        cout << a[n - 1] << endl;
    }
    for (int i = 0; i < n - 2; i++) {
        cin >> c[i];
    }
    sort(c, c + n - 2);
    ok = 1;
    for (int i = 0; i < n - 2; i++) {
        if (b[i] != c[i]) {
            cout << b[i] << endl;
            ok = 0;
            break;
        }
    }
    if (ok == 1) {
        cout << b[n - 2] << endl;
    }

    return 0;
}