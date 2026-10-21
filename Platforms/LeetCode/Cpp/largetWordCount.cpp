class Solution {
public:
    int countWords(const string& s) {
        int words = 1;

        for (char c : s) {
            if (c == ' ') {
                words++;
            }
        }

        return words;
    }

    string largestWordCount(vector<string>& messages, vector<string>& senders) {
        unordered_map<string, int> wordCount;

        for (int i = 0; i < messages.size(); i++) {
            wordCount[senders[i]] += countWords(messages[i]);
        }

        string ans = "";
        int maxWords = 0;

        for (auto& [sender, words] : wordCount) {
            if (words > maxWords ||
                (words == maxWords && sender > ans)) {

                maxWords = words;
                ans = sender;
            }
        }

        return ans;
    }
};