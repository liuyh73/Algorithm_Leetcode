## 题目

Given *n* non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.

![img](http://www.leetcode.com/static/images/problemset/rainwatertrap.png)
The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. **Thanks Marcos** for contributing this image!

**Example:**

```
Input: [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
```

 ## 思路与解法

利用两个指针：头指针，尾指针，头指针尾指针初始值作为两边边界高度。当头指针高度小于尾指针高度时，头指针向后移动，并比较当前位置高度与移动之前头指针高度大小，若前者小于后者，则蓄水量加上两者差值；若后者小于前者，则更新左边边界高度。当头指针高度大于或者等于尾指针高度时，尾指针向前移动，并比较当前位置高度与移动之前尾指针高度大小，若前者小于后者，则蓄水量加上两者差值；若后者小于前者，则更新右边边界高度。直到头指针等于尾指针，计算结束。

## 代码实现

```C
int trap(int* height, int heightSize) {
    // when the array is empty, we need to return zero.
    if(heightSize == 0) return 0;
    int *head=height, *tail=height+(heightSize-1);
    // minnEdge denotes the left or right edge, and is initialized with -1.
    // minnValue denotes the min value of left and right edge.
    int unitsOfRain=0, minnEdge=-1, minnValue;
    while(head != tail){
        if(*head < *tail){
            // when the minnEdge changes, the minnValue will change
            if(minnEdge == 1 || minnEdge == -1){
                minnValue=*head;
                minnEdge = 0;
            }
            head++;
            if(*head < minnValue) 
                unitsOfRain += minnValue - *head;
            // if the *head >= minnValue, the minnValue will change
            else
                minnValue = *head;
        }
        // the tail pointer is the same as head pointer
        else {
            if(minnEdge == 0 || minnEdge == -1){
                minnValue = *tail;
                minnEdge = 1;
            }
            tail--;
            if(*tail < minnValue)
                unitsOfRain += minnValue - *tail;
            else
                minnValue = *tail;
        }
    }
    return unitsOfRain;
}
```

## 遇到的问题及解决方法

最初代码中并未有数组为空时的判断处理，导致head指向null，tail计算出错：

```
Run Code Status: Runtime Error
```

加上判断条件即可完善代码：

```C
if(heightSize == 0) return 0;
```

