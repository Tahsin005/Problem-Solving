class Solution {
public:
    vector<int> findSubtreeSizes(vector<int>& parent, string s) {
        int n = parent.size();
        vector<vector<int>> children(n);
        for (int i = 1; i < n; i++) {
            children[parent[i]].push_back(i);
        }

        vector<int> last(26, -1);
        vector<int> saved(n, -1);
        vector<int> size(n, 1);

        vector<int> st;
        st.push_back(0);

        while (!st.empty()) {
            int x = st.back();
            st.pop_back();

            if (x >= 0) {
                int c = s[x] - 'a';
                saved[x] = last[c];
                last[c] = x;
                st.push_back(~x);
                for (int ch : children[x]) {
                    st.push_back(ch);
                }
            } else {
                x = -x - 1;
                int c = s[x] - 'a';
                last[c] = saved[x];
                if (x != 0) {
                    int newParent = (saved[x] != -1) ? saved[x] : parent[x];
                    size[newParent] += size[x];
                }
            }
        }
        return size;
    }
};
