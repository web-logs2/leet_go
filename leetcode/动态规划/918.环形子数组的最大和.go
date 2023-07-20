package 动态规划

/**
题目：给定一个首尾相连的数组，其实就是循环数组，找到所有子数组中最大和？

	输入：nums = [5,-3,5]
	输出：10
	解释：从子数组 [5,5] 得到最大和 5 + 5 = 10
*/

/**
方案一：单调递增队列
本题为「53. 最大子数组和」的进阶版 => 最大的子数组分为2种情况
	1、在nums[i:j]中间一段
	2、由2段组成 nums[0:i] + nums[j:n]
所以我们，拼接2个nums, 切且要求所求的最大和的子数组长度<=n， 也就是 nums[j:i],  要求 i-n<=j<i

T = O(n)
S = O(n)
*/
func maxSubarraySumCircular(nums []int) int {
	type pair struct {
		idx int
		val int
	}
	n := len(nums)
	pre, res := nums[0], nums[0]
	q := []pair{{0, pre}}
	for i := 1; i < 2*n; i++ {
		for len(q) > 0 && q[0].idx < i-n { //踢出步骤一：nums[j:i]长度超n了，丢弃
			q = q[1:]
		}
		//我们令pre = 截止到i的前面所有元素的和
		pre += nums[i%n]
		res = max(res, pre-q[0].val)               //计算 nums[j:i]的最大和
		for len(q) > 0 && q[len(q)-1].val >= pre { //踢出步骤二：保证队列的单调性， 作为【被减数】大于【减数】，本就没有意义
			q = q[:len(q)-1]
		}
		q = append(q, pair{i, pre})
	}
	return res
}
