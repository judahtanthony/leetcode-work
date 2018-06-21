#include <iostream>
using namespace std;

class Solution {
public:
    bool isPalindrome(int x) {
        if (x < 0) return false;
        if (x < 10) return true;
        // Find the end.
        int end = 10;
        while (x/end >= 10) end *= 10;
        while (true) {
            if (x%10 != (x/end)%10) return false;
            if (x < 10 || end < 100) break;
            x /= 10;
            end /= 100;
        }
        return true;
    }
};

int main(int argc, char** argv) {
    Solution s;
    cout << (s.isPalindrome(121) ? "true" : "false") << '\n';
    return EXIT_SUCCESS;
}
