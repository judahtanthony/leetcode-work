#include <iostream>
#include <stack>

using namespace std;

class Solution {
public:
    bool isValid(string s) {
        if (s.size() == 0) return true;
        if (s.size() % 2) return false;
        
        stack<char> seen;
        for (char c : s) {
            switch (c) {
                case '(':
                case '[':
                case '{':
                    seen.push(c);
                    break;
                case ')':
                case ']':
                case '}':
                    if (seen.empty() || (c - seen.top()) & 0xfc) {
                        return false;
                    }
                    seen.pop();
                    break;
                default:
                    return false;
            }
        }
        return seen.empty();
    }
};

int main() {
    Solution s;
    cout << s.isValid("(([){]})") << " should be " << 0 << endl;
    cout << s.isValid("()") << " should be " << 1 << endl;
    cout << s.isValid("()[]{}") << " should be " << 1 << endl;
    cout << s.isValid("(]") << " should be " << 0 << endl;
    cout << s.isValid("([)]") << " should be " << 0 << endl;
    cout << s.isValid("{[]}") << " should be " << 1 << endl;
}