package Alogrithm1

//快速排序算法
//算法描述：是对插入算法的一种优化，利用对问题的二分化，实现递归完成快速排序 ，在所有算法中二分化是最常用的方式，将问题尽量的分成两种情况加以分析， 最终以形成类似树的方式加以利用，因为在比较模型中的算法中，最快的排序时间 负载度为 O(nlgn).
//
//算法步骤
//将数据根据一个值按照大小分成左右两边，左边小于此值，右边大于
//将两边数据进行递归调用步骤1
//将所有数据合并
func QuickSort(arr []int) []int {
	if len(arr)<=1{
		return arr
	}
	splitData:=arr[0]
	low:=make([]int,0,0)//小
	height:=make([]int,0,0)//大
	mid:=make([]int,0,0)//一样大
	for i := 0; i < len(arr); i++ {
		if arr[i]<splitData {
			low=append(low,arr[i])
		}else if arr[i]>splitData{
			height=append(height,arr[i])
		}else {
			mid=append(mid,arr[i])
		}
	}
	low=QuickSort(low)
	height=QuickSort(height)
	myarr:=append(append(low,mid...),height...)
	return myarr
}


