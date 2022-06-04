# C#视频进度 
https://www.youtube.com/watch?v=GhQdlIFylQ8&list=LL&index=1&t=231s (看完了)                                
# Leetcode算法
### 搜索二维矩阵
题目：        
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：        
每行中的整数从左到右按升序排列。        
每行的第一个整数大于前一行的最后一个整数。        
        
思路： //矩阵特征是每行从小到大，且matrix [0]那一列也是从小到大，根据性质利用二分查找寻找target        
func searchMatrix(matrix [][]int, target int) bool {        
    row := sort.Search(len(matrix), func(i int) bool { return matrix[i][0] > target }) - 1  //用go中自带的sort.search找到目标函数的行数        
    if row < 0 {        
        return false        
    }        
    col := sort.SearchInts(matrix[row], target) //找对应的行中元素        
    return col < len(matrix[row]) && matrix[row][col] == target        
}        
        
### 搜索旋转排序数组
题目：        
整数数组 nums 按升序排列，数组中的值 互不相同 。        
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。        
给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。        
        
思路：        //也是二分查找，只不过要在二分查找的基础上先判断mid这个中间值是否比nums[0]要大，是的话就说明没旋转过半，否的话就说明旋转过半，再根据这个性质再进行二分查找         
func search(nums []int, target int) int {        
    left, right := 0, len(nums)-1        
    for left <= right {        
        mid := (right - left) / 2 + left        
        if nums[mid] == target {        
            return mid        
        }        
        if nums[mid] >= nums[left] {        
            if nums[mid] > target && target >= nums[left] {        
                right = mid - 1        
            } else {        
                left = mid + 1        
            }        
        } else {        
            if nums[mid] < target && target <= nums[right] {        
                left = mid + 1        
            } else {        
                right = mid - 1        
            }        
        }        
    }        
    return -1        
}        
        