#include<bits/stdc++.h>
#define ll long long
using namespace std;
/*



*/

int main() {
	int n; cin >> n;
	
	vector<int> v(n);
	
	for (int i = 0; i < n; i++) cin >> v[i];

	int sum = 0;

	for (int val : v) sum += val;

	cout << val << endl;
	


	return 0;
}
