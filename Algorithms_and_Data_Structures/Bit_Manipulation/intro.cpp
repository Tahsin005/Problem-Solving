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
#define Tahsin ios_base::sync_with_stdio(false), cin.tie(0),cout.tie(0);
using namespace std;
int main(){
    int a,b;
    cin>>a>>b;
    cout<<a + b<<el;

    // Find kth bit
    // int k;
    // cin>>k;
    // if((b>>k) & 1){
    //     cout<<1<<el;
    // }
    // else{
    //     cout<<0<<el;
    // }

    // FANCY WAY OF DOING ADDITION
    // a + b = (a ^ b ) + 2 * (a & b)
    // a + b = (a | b ) + (a & b)
    cout<<(a ^ b) + 2 * (a & b)<<el;
    cout<<(a | b ) + (a & b)<<el;

    
    return 0;    
}