#include <iostream>
#include <map>

using namespace std;

class Solution {
public:
    Solution() {
        index = {
            {'I', 1},
            {'V', 5},
            {'X', 10},
            {'L', 50},
            {'C', 100},
            {'D', 500},
            {'M', 1000}
        };
    }
    int romanToInt(string s) {
        if (s.empty()) {
            return 0;
        }
        int n = s.size();
        if (n == 1) {
            return index[s[0]];
        }
        int ret = index[s[0]];
        for  (int i = 1; i < n; ++i) {
            ret += index[s[i]];
            if (index[s[i-1]] < index[s[i]]) {
                ret -= index[s[i-1]] * 2;
            }
        }
        return ret;
    }
private:
    map<char, int> index;
};

int main() {
    Solution s;
    cout << s.romanToInt("III") << endl;
    cout << s.romanToInt("IV") << endl;
    cout << s.romanToInt("MCMXCVII") << endl;
    cout << s.romanToInt("XLIX") << endl;
}
