class Solution {
    public int reverse(int x) {
        int neg = 1;
        if (x == 0) {
            return 0;
        }
        if (x < 0) {
            neg *= -1;
            x *= -1;
        }
        int r = 0;
        while (x > 0) {
            // Overflow check.
            if (r > Integer.MAX_VALUE / 10) {
                return 0;
            }
            r = (r * 10);
            int ones = x % 10;
            // Overflow check.
            if (r > Integer.MAX_VALUE - ones) {
                return 0;
            }
            r += ones;
            x /= 10;
        }
        return r * neg;
    }
}
