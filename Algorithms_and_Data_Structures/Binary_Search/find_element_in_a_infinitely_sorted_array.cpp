#include<bits/stdc++.h>
#define ll long long
#define yes cout<<"YES"<<endl;
#define no cout<<"NO"<<endl;
#define nl cout<<'\n';
#define el '\n'
#define vtr vector
#define vii vector<int> 
#define vll vector<ll>
#define imax INT_MAX
#define imin INT_MIN
#define pb push_back
#define pii pair<int,int>
#define out(x) cout<<x<<"\n";
#define out2(x,y) cout<<x<<" "<<y<<"\n";
#define messi cout<<"Que Miras Bobo!"<<'\n';
#define Tahsin ios_base::sync_with_stdio(false), cin.tie(0),cout.tie(0);
using namespace std;
int main(){
    vii v = {1,2,3,4,5,6,7,8,9,10,11,12,13}; // Suppose it as an infinite array

    int key = 9;

    int l = 0;
    int r = 1;

    while(key > v[r]){
        l = r;
        r *= 2;
    }

    
    // Now normal binary search
    int elem = key;
    bool flag = false;
    int mid = 0;
    while(l <= r){
        mid = l + (r - l) / 2;
        if(v[mid] == elem){
            flag = true;
            break;
        }
        else if(elem < v[mid]){
            r = mid - 1;
        }
        else{
            l = mid + 1;
        }
    }
    if(flag == true) cout<<mid<<endl;
    else cout<<"-1"<<endl;
    return 0;    
}