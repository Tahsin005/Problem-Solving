#include <bits/stdc++.h>
using namespace std;
#define ll long long
const int M = 1e18+7;

const int N = 1e7+10;
vector<bool> isPrime(N, 1);

void solve() {
    isPrime[0] = isPrime[1] = false;
    for (int i = 2; i < N; i++) {
        if (isPrime[i] == true) {
            for (int j = 2 * i; j < N; j += i) {
                isPrime[j] = false;
            }
        }
    }
    
    for (int i = 0; i < 100; i++) {
        cout << i << " " << (isPrime[i] ? "Prime Number": "Not a Prime Number") << endl;
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt = 1; while (tt--) solve();
}
