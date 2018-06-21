import java.util.HashMap;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        int[] ret = new int[2];
        HashMap<Integer, Integer> cache = new HashMap<Integer, Integer>();;
        for (int i = 0; i < nums.length; ++i) {
            if (cache.containsKey(nums[i])) {
                ret[0] = cache.get(nums[i]);
                ret[1] = i;
                break;
            }
            cache.put(target-nums[i], i)    ;
        }
        return ret;
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        int[] nums = {2,7,11,15};
        int[] offsets = s.twoSum(nums, 9);
        System.out.println(offsets[0] + ", " + offsets[1]);
    }
}
