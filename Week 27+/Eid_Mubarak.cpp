#include <bits/stdc++.h>
using namespace std;

void typeWriterEffect(const string& text, int delay = 50, int scrambleSpeed = 15) {
    const string alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";

    for (char target : text) {
        if (target == ' ') {
            cout << ' ' << flush;
            this_thread::sleep_for(chrono::milliseconds(delay * 2));
            continue;
        }

        for (char c : alphabet) {
            cout << c << "\b" << flush;
            this_thread::sleep_for(chrono::milliseconds(scrambleSpeed));
            if (c == target) break;
        }

        cout << target << flush;
        this_thread::sleep_for(chrono::milliseconds(delay));
    }
    cout << endl;
}

int main() {
    string text = "Eid Mubarak!";
    typeWriterEffect(text);
    return 0;
}
