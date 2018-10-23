## 题目

Given an unsorted integer array, find the smallest missing positive integer.

**Example 1:**

```
Input: [1,2,0]
Output: 3
```

**Example 2:**

```
Input: [3,4,-1,1]
Output: 2
```

**Example 3:**

```
Input: [7,8,9,11,12]
Output: 1
```

## 思路与解法

- 一般算法：

  可以采用双重循环，从1-N依次遍历，查找是否在输入数组中存在过，时间复杂度为$O(N^2)$；

- 排序：

  可以先对输入数组进行排序，归并、快排，然后从小到大遍历一遍即可得到结果，时间复杂度为$O(NlogN)$；

- 桶：

  假如数据规模并不大，则可以利用桶，对数组中出现过的数字进行标记，然后从小到大遍历一遍即可得到结果，时间复杂度为O(1)，空间复杂度为O(N)。

但是，题目中要求时间复杂度为O(N)，空间复杂度为O(1)，所以上述三种方法都不成立，我们可以采用以下思路：

我们需要覆盖原有的数组，可以将1放在0号位置上，将2放在1号位置上……将N放在N-1号位置上，此时，对于一个恰好包含0-N的长度为N的数组，当我们将每个数字归位后，会发现，数组中的每一个数字都应该满足`nums[i]=i+1`或者`nums[nums[i]-1]=nums[i]`。利用这种思路，我们会得到以下代码：

## 代码实现

```C
// 交换两数字
void swap(int *a, int *b){
    int temp = *a;
    *a = *b;
    *b = temp;
}

int firstMissingPositive(int* nums, int numsSize) {
    // 遍历数组中的每一个数字
    for(int i=0;i<numsSize;i++){
        // 当数字不在1-numsSize范围内时，直接跳过
        if(nums[i]<0 || nums[i]>numsSize)
            continue;
        // 当数字在1-numsSize范围内，判断数字是否满足nums[nums[i]-1] != nums[i]
        // 若满足该条件，则交换nums[nums[i]-1]和nums[i]，使得在i位置上的数字归位
        // 并且将i减1，使得在下次循环时，继续判断当前的第i位数字。
        else if(nums[nums[i]-1] != nums[i]){
            swap(&nums[nums[i]-1], &nums[i]);
            i--;
        }
    }
    // 交换处理完成之后，便利结果判断the first missing positive
    for(int i=0;i<numsSize;i++)
        if(nums[i] != i+1)
            return i+1;
    // 若没有丢失正整数，则返回numsSize+1
    return numsSize+1;
}
```

## 遇到的问题及解决方法

最初在循环中交换条件写做如下：

```c
if(nums[i] != i+1){
	swap(&nums[nums[i]-1], &nums[i]);
	i--;
}
```

上述判断条件是根据当前位置判断当前数字，结果发现当数组中出现相同的正整数时会陷入死循环，比如：[-1,3,3,4,5]。当i等于1时，nums[1] != 1+1，所以swap(&nums[2],&nums[1])，交换之后对原数组无影响，并且i进行减一操作，陷入死循环。

解决方法：

```c
if(nums[nums[i]-1] != nums[i]){
	swap(&nums[nums[i]-1], &nums[i]);
	i--;
}
```

对于上述判断条件，则是根据当前数字决定其位置，避免了陷入死循环。