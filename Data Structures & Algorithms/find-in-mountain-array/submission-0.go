/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
	n:=mountainArr.length()
	left,right:=1,n-1

	for left<right {
		mid:=(left+right)/2

		if mountainArr.get(mid)<mountainArr.get(mid+1) {
			left=mid+1
		} else {
			right=mid
		}
	}
	peak:=left

	left,right=0,peak
	
	for left<=right {
		mid:=(left+right)/2

		if mountainArr.get(mid)==target {
			return mid 
		} else if target>mountainArr.get(mid) {
			left=mid+1
		} else {
			right=mid-1
		}
	}

	left,right=peak+1,n-1

	for left<=right {
		mid:=(left+right)/2

		if mountainArr.get(mid)==target {
			return mid 
		} else if target<mountainArr.get(mid) {
			left=mid+1
		} else {
			right=mid-1
		}
	}
	return -1
}
