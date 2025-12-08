#include <bits/stdc++.h>
using namespace std;
#define ll long long
const int M = 1e18+7;

bool is_prime(int n) {
    if (n == 0 or n == 1) return false;
    for (int i = 2; i * i <= n; i++) {
        if (n % i == 0) {
            return false;
        }
    }
    
    return true;
}

vector<int> prime_factors(int n) {
    vector<int> ans;
    if (n == 0 or n == 1) return ans;
    
    for (int i = 2; i * i <= n; i++) {
        while (n % i == 0) {
            ans.push_back(i);
            n /= i;
        }
    }
    
    if (n > 1) {
        ans.push_back(n);
    }
    
    return ans;
}

void solve() {
    cout << is_prime(1) << " " << is_prime(2) << " " << is_prime(4) << endl;
    
    for (int prime: prime_factors(24)) {
        cout << prime << endl;
    }
    cout << endl;
    
    for (int prime: prime_factors(36)) {
        cout << prime << endl;
    }
    cout << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt = 1; while (tt--) solve();
}
