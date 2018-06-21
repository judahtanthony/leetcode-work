#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int> ret;
        unordered_map<int, int> cache;
        for (int i = 0; i < nums.size(); ++i) {
            if (cache.find(nums[i]) != cache.end()) {
                ret.push_back(cache[nums[i]]);
                ret.push_back(i);
                break;
            }
            cache[target-nums[i]] = i;
        }
        return ret;
    }
};

int main() {
    vector<int> nums = { 2, 7, 11, 15 };
    Solution s;
    vector<int> offsets = s.twoSum(nums, 9);
    cout << offsets[0] << ", " << offsets[1] << '\n';

    return EXIT_SUCCESS;
}
