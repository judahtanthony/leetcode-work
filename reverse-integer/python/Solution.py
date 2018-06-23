INT_MAX = (1 << 31) - 1;
class Solution:
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        neg = 1;
        if x == 0:
            return 0
        
        if x < 0:
            neg *= -1
            x *= -1
            
        r = 0
        while x > 0:
            # Overflow check.
            if r > INT_MAX / 10:
                return 0
            r *= 10
            ones = x % 10
            # Overflow check.
            if r > INT_MAX - ones:
                return 0
            r += ones
            x = math.floor(x / 10)
        
        return r * neg
    
