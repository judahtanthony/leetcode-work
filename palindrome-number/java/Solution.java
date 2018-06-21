class Solution {
    public boolean isPalindrome(int x) {
        if (x < 0) return false;
        if (x < 10) return true;
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

    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.isPalindrome(121) ? "true" : "false");
    }
}
