#include<bits/stdc++.h>
using namespace std;
void TEST_CASE() {
    int k;
    cin >> k;
    if (k % 3 == 1) {
        cout << "Yes\n";
    } else {
        cout << "No\n";
    }
}
int main() {
    ios::sync_with_stdio(0), cin.tie(0);
    int tt = 1;
    cin >> tt;
    while (tt--) {
        TEST_CASE();
    }
}