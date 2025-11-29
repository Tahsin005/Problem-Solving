#include <iostream>
using namespace std;
int main() {
    int t;
    long long n;
    cin >> t;
    while (t--) {
        cin >> n;
        if (n == 1) cout << "-1" << endl;
        else {
            cout << "2";
            while (n != 1)
                cout << "3", n--;
            cout << endl;
        }
    }
    return 0;
}