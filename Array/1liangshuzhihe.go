package Array

func twoSum1(nums []int, target int) []int {
	res := []int{0, 1}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = []int{i, j}
			}
		}
	}
	return res
}

// 注意在map中存放的key为val val为index
func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for index, val := range nums {
		if _, ok := hashMap[target-val]; ok {
			return []int{index, hashMap[target-val]}
		}
		hashMap[val] = index
	}
	return nil
}

func main() {

}
