//                      ...              ...                                                             
//    .::::::::.       .@@+              %@%                                                             
//    =@@@@@@@@+       .@@+              +*+                                                             
//    -**#@@%**-       .@@+                                                                              
//       :@@-    :=+:  .@@+.==    .=+-   ::: ::: -+:                                                     
//       :@@-   #@@@@% .@@#@@@%  +@@@@@: %@% #@##@@@-                                                    
//       :@@-  +@@+*@@+.@@@+%@@=.@@+:%@% %@% #@@**@@%                                                    
//       :@@-  :--  @@*.@@+ .@@*:@@= ::: %@% #@%  *@@                                                    
//       :@@-   -*##@@*.@@+  @@* %@@@*-  %@% #@%  *@@                                                    
//       :@@-  =@@#*@@*.@@+  @@*  +#@@@* %@% #@%  *@@                                                    
//       :@@-  @@#  @@*.@@+  @@*:== .#@@ %@% #@%  *@@                                                    
//       :@@-  %@@=*@@#.@@+  @@*-@@=:%@@ %@% #@%  *@@                                                    
//       :@@-  -@@@%%@@.@@+  @@* *@@@@@- %@% #@%  *@@                                                    
//       .--.   :+= :-- --.  :-:  :=+-.  :-: :-:  :--               

#include<bits/stdc++.h>
#define ll long long
#define cyes cout << "YES" << '\n';
#define yes cout << "Yes" << '\n';
#define cno cout << "NO"<< '\n';
#define no cout << "No"<< '\n';
#define minus cout << -1 << '\n';
#define nl cout << '\n';
#define el '\n'
#define all(x) x.begin(), x.end()
#define vtr vector
#define vii vector<int> 
#define vll vector<ll>
#define imax INT_MAX
#define imin INT_MIN
#define pb push_back
#define pii pair<int,int>
#define out(x) cout << x << '\n';
#define out2(x,y) cout << x << " " << y << '\n';
#define messi cout << "Que Miras Bobo!" << '\n';
#define e4 int main()
#define Tahsin ios_base::sync_with_stdio(false), cin.tie(0),cout.tie(0);
using namespace std;
//@ Problem Link: 
void solve(){
    int n,x,k;
    cin >> n >> x >> k;
    int boys = x;
    int girls = n - x;

    if(boys % k == 0 and girls % k == 0){
        cout << 0 << el;
        return;
    }

    int remboys = boys % k;
    int remgirls = girls % k;

    if(remboys == remgirls){
        cout << 0 << el;
        return;
    }

    cout << abs(remboys - remgirls) << el;

    
}
e4{
    Tahsin
    ll t = 1;
    cin >> t;
    for(ll cs = 1; cs <= t; cs++){
        // cout << "Case #" << cs << ": ";
        solve();
    }
    return 0;    
}