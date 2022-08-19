package ac

func nextPermutation(nums []int) {
    i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
    for i >= 0 && nums[i] >= nums[j] {
        i--
        j--
    }

    if i >= 0 {
        for nums[k] <= nums[i] {
            k--
        }
        nums[i], nums[k] = nums[k], nums[i]
    }

    for i,j:=j,len(nums)-1;i<j;i,j=i+1,j-1{
        nums[i],nums[j]=nums[j],nums[i]
    }

}
