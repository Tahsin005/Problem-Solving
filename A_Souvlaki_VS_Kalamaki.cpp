#include <bits/stdc++.h>
using namespace std;
const int N = 102;
int a[N];
void TEST_CASE() {
    int n;
    cin >> n;
    for (int i = 1; i <= n; i++) cin >> a[i];
    sort(a + 1, a + n + 1);
    for (int i = 2; i < n; i += 2) {
        if (a[i] != a[i + 1]) {
            cout << "NO" << "\n";
            return;
        }
    }
    cout << "YES" << "\n";
}
int main() {
    int tt;
    cin >> tt;
    while (tt--) TEST_CASE();
    return 0;
}