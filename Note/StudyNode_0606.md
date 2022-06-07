# Leetcode算法
### 快速排序
题目：用go实现快速排序          
思路：         //从最后一个开始进行一遍快排，再以第一遍顺序为基准左右分开再快排，节省时间复杂度，分治排序
func QuickSort(nums []int) []int {         
    // 思路：把一个数组分为左右两段，左段小于右段         
    quickSort(nums, 0, len(nums)-1)         
    return nums         
         
}         
// 原地交换，所以传入交换索引         
func quickSort(nums []int, start, end int) {         
    if start < end {         
        // 分治法：divide         
        pivot := partition(nums, start, end)         
        quickSort(nums, 0, pivot-1)         //以end的新位置为基准左右再进行快排
        quickSort(nums, pivot+1, end)         
    }         
}         
// 分区         
func partition(nums []int, start, end int) int {         
    // 选取最后一个元素作为基准pivot         
    p := nums[end]         
    i := start         
    // 最后一个值就是基准所以不用比较         
    for j := start; j < end; j++ {         
        if nums[j] < p {         //算end的位置
            swap(nums, i, j)         
            i++         
        }         
    }         
    // 把基准值换到中间         
    swap(nums, i, end)         
    return i         
}         
// 交换两个元素         
func swap(nums []int, i, j int) {         
    t := nums[i]         
    nums[i] = nums[j]         
    nums[j] = t         
}         
### 归并排序         
题目：用go实现并归排序         
思路：         //分治法把数组不断细分，再排序后合并            
func MergeSort(nums []int) []int {         
    return mergeSort(nums)         
}         
func mergeSort(nums []int) []int {         
    if len(nums) <= 1 {         
        return nums         
    }         
    // 分治法：divide 分为两段         
    mid := len(nums) / 2         
    left := mergeSort(nums[:mid])         
    right := mergeSort(nums[mid:])         
    // 合并两段数据         
    result := merge(left, right)         
    return result         
}         
func merge(left, right []int) (result []int) {         
    // 两边数组合并游标         
    l := 0         
    r := 0         
    // 注意不能越界         
    for l < len(left) && r < len(right) {         
        // 谁小合并谁         
        if left[l] > right[r] {         
            result = append(result, right[r])         
            r++         
        } else {         
            result = append(result, left[l])         
            l++         
        }         
    }                  
    // 剩余部分合并         
    result = append(result, left[l:]...)         
    result = append(result, right[r:]...)         
    return         
}         
