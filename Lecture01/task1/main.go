func twoSum(nums []int, target int) []int {
	nmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		rNum := target - nums[i]
		reqNumIdx, isPresent := nmap[rNum]

		if isPresent {
			return []int{i, reqNumIdx}
		}
		nmap[nums[i]] = i
	}
	return []int{}
}