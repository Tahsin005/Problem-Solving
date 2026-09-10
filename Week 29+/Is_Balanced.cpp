#include <bits/stdc++.h>

using namespace std;

string ltrim(const string &);
string rtrim(const string &);
 
bool isSameType(char first, char second) {
    return (first == '{' and second == '}') or (first == '[' and second == ']') or (first == '(' and second == ')');
}

bool isOpen(char ch) {
    return ch == '{' or ch == '[' or ch == '(';
}

bool isClose(char ch) {
    return ch == '}' or ch == ']' or ch == ')';
}

string isBalanced(string s) {
    stack<char> st;
    
    for (int i = 0; i < s.length(); i++) {
        if (isOpen(s[i])) {
            st.push(s[i]);
        } else {
            if (!st.empty()) {
                char top_elem = st.top();
                if (isSameType(top_elem, s[i])) {
                    st.pop();
                } else {
                    return "NO";
                }
            } else {
                return "NO";
            }
        }
    }
    
    return st.empty()? "YES" : "NO";
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string t_temp;
    getline(cin, t_temp);

    int t = stoi(ltrim(rtrim(t_temp)));

    for (int t_itr = 0; t_itr < t; t_itr++) {
        string s;
        getline(cin, s);

        string result = isBalanced(s);

        fout << result << "\n";
    }

    fout.close();

    return 0;
}

string ltrim(const string &str) {
    string s(str);

    s.erase(
        s.begin(),
        find_if(s.begin(), s.end(), not1(ptr_fun<int, int>(isspace)))
    );

    return s;
}

string rtrim(const string &str) {
    string s(str);

    s.erase(
        find_if(s.rbegin(), s.rend(), not1(ptr_fun<int, int>(isspace))).base(),
        s.end()
    );

    return s;
}
