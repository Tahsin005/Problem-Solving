vector<int> intersectSorted(const vector<int>& A, const vector<int>& B) {
    int i = 0, j = 0;
    vector<int> result;

    while (i < A.size() && j < B.size()) {
        if (A[i] == B[j]) {
            result.push_back(A[i]);
            i++;
            j++;
        } else if (A[i] < B[j]) {
            i++;
        } else {
            j++;
        }
    }

    return result;
}
