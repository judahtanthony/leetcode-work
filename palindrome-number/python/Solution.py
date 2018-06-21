import math

class Solution:
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        if x < 0:
            return False
        if x < 10:
            return True
        end = 10
        while x/end >= 10:
            end *= 10
        while True:
            if x%10 != math.floor(x/end)%10:
                return False
            if x < 10 or end < 100:
                break
            x = math.floor(x/10)
            end = math.floor(end/100)

        return True

if __name__ == "__main__":
    s = Solution()
    print("true" if s.isPalindrome(121) else "false")
