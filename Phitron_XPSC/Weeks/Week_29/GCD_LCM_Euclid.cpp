#include <bits/stdc++.h>
using namespace std;
#define ll long long


/*

GCD and LCM using Euclid's Algorithm

*/

int GCD(int a, int b) {
    if (b == 0) return a;
    return GCD(b, a % b);
}

int LCM(int a, int b) {
    return (a * b) / GCD(a, b);
}

void solve() {
    cout << GCD(4, 12) << endl;
    cout << GCD(18, 12) << endl;
    
    cout << LCM(4, 12) << endl;
    cout << LCM(12, 18) << endl;
    
    // minimum fraction (12/18    ->    2/3)
    // (a / gcd(a, b)) / (b / gcd(a, b))
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    solve();
}
