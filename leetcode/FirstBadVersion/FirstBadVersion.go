package FirstBadVersion

import "sort"

func firstBadVersion(n int) int {       
    return sort.Search(n, func(version int) bool { return isBadVersion(version) })        
    //golang中自带的二分查找函数 sort.Search(n int,f func(i int) bool) int 会从[0, n)中取出一个值index，index为[0, n)中最小的使函数f(index)为True的值              
} 


//没啥用
func isBadVersion(ver int) bool {
	return true
}
