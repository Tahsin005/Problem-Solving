#include <bits/stdc++.h>
using namespace std;
#define ll long long


/*

Given array a of n integers. All integers are present in the even count.
Except one. Find that one integer which has odd count in O(N) time complexity and 0(1) space;

N < 1e5;
a[i] < 1e5;

9
2 4 6 7 7 4 2 2 2
*/
void solve() {
    int n;
    cin >> n;
    
    vector<int> v(n);
    for (int i = 0; i < n; i++) {
        cin >> v[i];
    }
    
    int num = v[0];
    for (int i = 1; i < n; i++) {
        num ^= v[i];
    }
    
    cout << num << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    solve();
}
