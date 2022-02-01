// [1,2,3]
// [2,3,3]
// [3,4,3]

// 1,2,3,4
// 2,3,4,4
// 3,4,5,4
// 4,5,5,5
// 5,6,6,5
// 6,7,6,6
// 7,7,7,7

// -1,1
// 0,1
// 1,1

// while true
// for i=0 to n
// if num[i]+1 <= num[i+1] num[i]++

// public int minMoves(int[] nums) {
//         int min = nums[0];
//         for(int num : nums) min = Math.min(min, num);

//         int res = 0;
//         for(int num : nums) res += (num - min);

//         return res;
//     }

// sum + m * (n - 1) = x * n
// x = minNum + m
// sum - minNum * n = m

func minMoves(nums []int) int {
	//     equal := true
	//     for i:=0;i<len(nums)-1;i++ {
	//         if nums[i]!=nums[i+1] {
	//             equal = false
	//         }
	//     }
	//     if equal { return 0 }

	//     sort.Ints(nums)

	//     ans := 0
	//     for true {
	//         equal := true
	//         for i:=0;i<len(nums)-1;i++ {
	//             if nums[i] != nums[i+1] { equal = false }
	//             if i+1<len(nums) && nums[i] <= nums[i+1] {
	//                 nums[i] += 1
	//             } else {
	//                 nums[i+1] += 1
	//             }
	//         }

	//         // fmt.Println(nums)
	//         if equal { break }
	//         ans += 1
	//     }

	//     return ans

	sum := 0
	min := nums[0]
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] < min {
			min = nums[i]
		}
	}

	return sum - min*len(nums)
}
