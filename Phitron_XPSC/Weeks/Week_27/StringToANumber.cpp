#include <bits/stdc++.h>
# define int int64_t
# define e4 int32_t main
# define el "\n"
using namespace std;

map<string, int> wordToNum = {
    {"one", 1},
    {"two", 2},
    {"three", 3},
    {"four", 4},
    {"five", 5},
    {"six", 6},
    {"seven", 7},
    {"eight", 8},
    {"nine", 9},
    {"ten", 10},
    {"eleven", 11},
    {"twelve", 12},
    {"thirteen", 13},
    {"fourteen", 14},
    {"fifteen", 15},
    {"sixteen", 16},
    {"seventeen", 17},
    {"eighteen", 18},
    {"nineteen", 19},
    {"twenty", 20},
    {"thirty", 30},
    {"forty", 40},
    {"fifty", 50},
    {"sixty", 60},
    {"seventy", 70},
    {"eighty", 80},
    {"ninety", 90},
    {"hundred", 100},
    {"thousand", 1000},
    {"million", 1000000},
    {"billion", 1000000000}
};

int wordsToNumber(string &input) {
    stringstream ss(input);
    string word;

    vector<string> words;
    while (ss >> word) {
        for (auto &c : word) c = tolower(c);
        words.push_back(word);
    }

    int total = 0, current = 0;

    for (string &w : words) {
        if (wordToNum.find(w) != wordToNum.end()) {
            int num = wordToNum[w];

            if (num == 100) {
                current *= num;
            } else if (num >= 1000) {
                total += current * num;
                current = 0;
            } else {
                current += num;
            }
        }
    }

    return total + current;
}

e4 () {
    string input;
    getline(cin, input);
    
    cout << wordsToNumber(input) << el;

    return 0;
}
