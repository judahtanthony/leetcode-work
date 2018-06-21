#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        if (strs.size() == 0) return "";
        if (strs.size() == 1) return strs[0];
        string longest = strs[0];
        for (int j = 0; j < longest.size(); ++j) {
            for (int i = 1; i < strs.size(); ++i) {
                if (j >= strs[i].size() || longest[j] != strs[i][j]) {
                    return longest.substr(0, j);
                }
            }
        }
        return longest;
    }
};

int main() {
    Solution s;
    vector<string> strs = {
        "flower",
        "flow",
        "flight"
    };
    cout << s.longestCommonPrefix(strs) << endl;
    strs = {"dog","racecar","car"};
    cout << s.longestCommonPrefix(strs) << endl;
}
