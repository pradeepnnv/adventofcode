package day1

func CaliberationValuePart1(list []string) int {
	sum := 0
	for _, v := range list {
		nums := make([]int, 0)
		for _, c := range v {
			// fmt.Println(c)
			if c >= 48 && c <= 57 {
				nums = append(nums, int(c)-'0')
			}
		}
		if len(nums) == 1 {
			nums = append(nums, nums[0])
		}
		sum += nums[0]*10 + nums[len(nums)-1]
	}
	return sum
}
