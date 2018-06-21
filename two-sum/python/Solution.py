class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        cache = {}
        for k in range(len(nums)):
            if nums[k] in cache:
                return [ cache[nums[k]], k ]
            cache[target - nums[k]] = k

        return []

if __name__ == "__main__":
    s = Solution()
    print(s.twoSum([2,7,11,15], 9))