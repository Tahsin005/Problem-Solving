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
    vtr<char> v = {'a','c','f','h'};

    char key = 'c';
    int l = 0;
    int r = v.size() - 1;
    int mid = 0;
    char res = '#';

    while(l <= r){
        mid = l + (r - l) / 2;
        if(v[mid] == key){
            l = mid + 1;
        }
        if(v[mid] < key){
            l = mid + 1;
        }
        else if(v[mid] > key){
            res = v[mid];
            r = mid - 1;
        }
    }
    if(res == '#') cout<<"No greater element found"<<el;
    else cout<<res<<" found"<<el;
    return 0;    
} 