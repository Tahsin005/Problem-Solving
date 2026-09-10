#include <bits/stdc++.h>
using namespace std;
#define ll long long


/*

Subset generate: using bit masking

4
1 4 7 9

*/

vector<vector<int>> gen_subsets(vector<int>& nums) {
    int n = nums.size();
    int subset_cnt = (1<<n);
    
    vector<vector<int> > subsets;
    
    for (int mask = 0; mask < subset_cnt; mask++) {
        vector<int> subset;
        for (int i = 0; i < n; i++) {
            if ((mask & (1<<i)) != 0) {
                subset.push_back(nums[i]);
            }
        }
        
        subsets.push_back(subset);
    }
    
    return subsets;
}
void solve() {
    int n;
    cin >> n;
    
    vector<int> v(n);
    for (int i = 0; i < n; i++) {
        cin >> v[i];
    }
    
    vector<vector<int> > subsets = gen_subsets(v);
    
    for (auto subset: subsets) {
        for (auto value: subset) {
            cout << value << " ";
        }
        cout << endl;
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    solve();
}
