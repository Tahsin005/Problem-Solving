#include <bits/stdc++.h>
using namespace std;
#define ll long long
const int M = 1e18+7;

int arr[16];
 
bool sorted(int l, int r) {
    for (int i = l + 1; i <= r; i++) {
        if (arr[i] < arr[i - 1])return false;
    }
    return true;
}
 
int solve(int l, int r) {
    if (sorted(l, r)) return r - l + 1;
    int mid = (r + l) >> 1;
    return max(solve(l, mid), solve(mid + 1, r));
}
 
void SUIII() {
    int n;
    cin >> n;
    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }
    cout << solve(0, n - 1);
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt = 1; 
    while (tt--) {
        SUIII();
    }
    
    return 0;
}
