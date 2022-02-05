class Solution:
    def twoSum(self, nums, target):
        d = {}
        for i in range(0, len(nums)):
            n = nums[i]
            k = target-n
            if k in d:
                return [i, d[k]]
            d[n] = i
        return []
