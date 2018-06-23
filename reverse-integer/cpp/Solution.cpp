#include <limits.h>
class Solution {
public:
    int reverse(int x) {
        int n = 1;
        if (x == 0) {
            return 0;
        }
        if (x <= 0) {
            n *= -1;
            x *= -1;
        }
        int r = 0;
        while(x > 0) {
            // Overflow check.
            if (r > INT_MAX / 10) {
                return 0;
            }
            r = (r * 10);
            int ones = x % 10;
            // Overflow check.
            if (r > INT_MAX - ones) {
                return 0;
            }
            r += ones;
            x /= 10;
        }
        return r * n;
    }
};
