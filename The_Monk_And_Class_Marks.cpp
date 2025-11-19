#include <bits/stdc++.h>
#define ll long long
#define endl "\n"
using namespace std;

// https://www.hackerearth.com/problem/algorithm/the-monk-and-class-marks/

void solve(){
    int n;
    cin >> n;
    
    map<int, multiset<string> > marks;
    for (int i = 0; i < n; i++) {
        int mark; string name;
        cin >> name >> mark;
        marks[mark].insert(name);
    }
    
    auto curr_it = --marks.end();
    while (true) {
        auto &students = (curr_it->second);
        int curr_mark = (curr_it->first);
        for (auto student: students) {
            cout << student << " " << curr_mark << endl;
        }
        if (curr_it == marks.begin()) break;
        curr_it--;
    }
}

int main(){
    ios_base::sync_with_stdio(false); cin.tie(NULL);
    solve();
}
