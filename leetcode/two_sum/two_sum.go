package two_sum

func twoSum(nums []int, target int) []int {
	dict := make(map[int]interface{})
	for index,val := range nums {
		expect := target - val
		if dict[expect] != nil {
			return []int{dict[expect].(int),index}
		}
		dict[val] = index;
	}
	return  []int{}
}