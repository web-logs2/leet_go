package 动态规划

/**
题目：给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和

	输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
	输出：6
	解释：连续子数组 [4,-1,2,1] 的和最大，为 6

思路：动态规划
我们用f(i)表示以nums[i]为结尾的，子数组的最大和，那么我们只需要求出 max([f(0), f(1), ...., f(n-1)])
	其中 f(i) = max(f(i-1)+nums[i], nums[i])
T = O(n)
S = O(1)
*/
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
