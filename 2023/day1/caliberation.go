package day1

import "log"

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

func CaliberationValuePart2(list []string) int {
	sum := 0
	numMap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	for _, line := range list {
		nums := make([]int, 0)
		for i := 0; i < len(line); i++ {
			if c := line[i]; c >= 48 && c <= 57 {
				log.Printf("Found %d num at index %d. Adding %d to the number list\n", int(c)-'0', i, int(c)-'0')
				nums = append(nums, int(c)-'0')
			} else {
				for k, v := range numMap {
					// log.Printf("Searching for %s at index %d\n", k, i)
					//k==>"nine"
					// len(line) == > 8, i==> 4, len(k)-1==>3
					// log.Println("Comparing for ", len(line)-i-len(k)-1)
					// log.Println("Comparing with ", line[i:len(k)])
					// if len(line)-i-len(k)-1 > 0 {
					// 	log.Println("String part is :" + line[i:len(k)])
					// }
					// log.Println("comp is ", len(line)-i-len(k))
					// log.Println("comparing with ", line[i:len(k)])
					if i+len(k) <= len(line) && line[i:len(k)+i] == k {
						nums = append(nums, v)
						// i += len(k) - 1

						log.Printf("Found %s at index %d. Adding %d to the number list\n", k, i, v)
						break
					}
				}
			}
			// log.Printf("Numbers found so far in %s is %v\n", line, nums)
		}
		// if len(nums) == 1 {
		// 	nums = append(nums, nums[0])
		// }
		// log.Printf("Numbers found so far in %s is %v\n", line, nums)
		sum += nums[0]*10 + nums[len(nums)-1]
		log.Printf("Sum after processing %s is %d\n", line, sum)
	}

	return sum
}
