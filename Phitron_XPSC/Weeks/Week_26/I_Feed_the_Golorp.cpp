#include <bits/stdc++.h>
using namespace std;

vector<int> v[1010];
int val[1010];

int main() {
	string s;
	cin >> s;
	int len = s.length(), pos;
	for (int i = 0; i < len; i++) {
		if (s[i] == ':' and s[i + 1] == '-') {
			pos = i;
			break;
		}
    }

	vector<int> seq;
	for(int i = 0; i < pos;) {
		if (s[i] != '_') {
			i++;
			continue;
		}
		int j = i;
		while( s[j] == '_')
			j++;
		int p = j - i;
		seq.push_back(p);
		i = j;
	}

	for(int i = pos + 1; i < len;) {
		if (s[i] != '_') {
			i++;
			continue;
		}
		int j = i;
		while (s[j] == '_')
			j++;
		int p1 = j - i;
		bool flag = (s[j] == '>');
		i = ++j;
		while (s[j] == '_')
			j++;
		int p2 = j - i;
		i = j;
		if (flag) v[p1].push_back(p2);
		else v[p2].push_back(p1);
	}

	while (true) {
		bool flag = true;
		for (int i : seq)
			for (int j : v[i])
				if (val[j] >= val[i]) {
					val[i]++;
					if (val[i] >= 10) {
						cout << "false";
						return 0;
					}
					flag = false;
				}
		if (flag)
			break;
	}

	for (int i : seq)
		cout << val[i];
	return 0;
}