#include <bits/stdc++.h>

#define ll long long
#define endl "\n"
using namespace std;

vector<int> nextGreaterElement(vector<int> &arr) {
    stack<int> stk;
    vector<int> nge(arr.size());
    
    for (int i = 0; i < arr.size(); i++) {
        while (!stk.empty() and arr[i] > arr[stk.top()]) {
            nge[stk.top()] = i;
            stk.pop();
        }
        stk.push(i);
    }
    
    while (!stk.empty()) {
        nge[stk.top()] = -1;
        stk.pop();
    }
    
    return nge;
}

void solve() {
    int n;
    cin >> n;
    
    vector<int> arr(n);
    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }
    
    vector<int> nge = nextGreaterElement(arr);
    for (int i = 0; i < n; i++) {
        cout << arr[i] << " " << (nge[i] == -1? -1: arr[nge[i]]) << endl;
    }
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    solve();
}
