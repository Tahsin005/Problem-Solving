#include<bits/stdc++.h>
#define ll long long
using namespace std;
/*



*/

int main() {
    int num;
    cin >> num;
    int arr[] = {4, 7, 44, 77, 47, 74, 777, 444, 747, 474, 744, 477, 774, 447};
    for (int val : arr) {
        if ( num % val == 0) {
            cout << "YES";
            return 0;
        }
    }

    cout << "NO";

    return 0;
}