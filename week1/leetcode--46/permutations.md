# leetcode--46.Permutations题解

##题目
Given a collection of distinct integers, return all possible permutations.
Example:
```
Input:[1,2,3]
Output:
[
	[1,2,3],
	[2,3,2],
	[2,1,3],
	[2,3,1],
	[3,1,2],
	[3,2,1]
]
```
##思路与解法
题目要求给定一个输入数组，输出该数组的全排列结果。

    1. 以[1,2]数组为例： 主要思路：

     - 将第一个数字1放到最前面，然后对数字2进行全排列，所以输出[1,2]
     - 将第二个数字2放在最前面(与数字1交换)，然后对数字1进行全排列，所以输出[2,1]
   2. 扩大数组的规模，以[1,2,3]为例：
     - 将第一个数字1放在最前面，准备输出[1,X,X]，然后对2，3进行全排列即为后两位
     - 将第二个数字2放在最前面，准备输出[2,X,X]，然后对1，3进行全排列即为后两位
     - 将第三个数字3放在最前面，准备输出[3,X,X]，然后对1，2进行全排列即位后两位
   3. 以此类推，当数组中由N个数：
      - 将第一个数字$D_1$放在最前面，准备输出[$D_1$,X,X……]，然后对后面N-1个数进行全排列
      - 将第二个数字$D_2$放在最前面，准备输出[$D_2$,X,X……]，然后对后面N-1个数进行全排列
      - ……
      - 将第N个数字$D_N$放在最前面，准备输出[$D_N$,X,X……]，然后对后面N-1个数进行全排列


由上述分析可知，这是一个递归过程，把对整个序列做全排列的问题归结为了对其子序列做全排列；在实际操作中，可以通过交换swap操作来讲后面的数字放在前面。

## 代码实现

题目中给定了缺失的函数：

```C
/**
 * Return an array of arrays of size *returnSize.
 * Note: The returned array must be malloced, assume caller calls free().
 */
 int** permute(int* nums, int numsSize, int* returnSize) {
 
 }
```

在该函数中，原数组规模为numsSize，原数组指针为nums，全排列得到的数组数量指针为returnSize，返回值为全排列结果的二维数组的指针。具体实现如下：

swap函数用于数据交换，便于将后续数字放置到前面；

handlePermute函数用于得到全排列的数组，参数分别为原数组指针nums，某个全排列的开始下标begin，原数组的规模numsSize，全排列得到的所有结果的二维数组permutedNums，全排列结果数量指针returnSize；

permute函数则进行内存分配等一些初始化和结果处理工作。

```C
void swap(int *a, int *b){
    int temp = *a;
    *a = *b;
    *b= temp;
}
void handlePermute(int *nums, int begin, int numsSize, int *permutedNums[], int *returnSize){
    // begin>=numsSize表示已经得到一个全排列，没有子全排列需要再次进行了
    if(begin>=numsSize){
        // 存储该全排列结果
        for(int i=0;i<numsSize; i++) {
            permutedNums[*returnSize][i]=nums[i];
        }
        // 计数加一
        (*returnSize)++;
        return;
    }
    // 该循环，用于计算全排列
    for(int i=begin; i<numsSize; i++) {
        // 该条件句，表示将后续数据nums[i]利用swap提到第一位
        if(i > begin)
            swap(&nums[begin], &nums[i]);
        handlePermute(nums, begin+1, numsSize, permutedNums, returnSize);
        // 回溯过程中，利用swap将nums[i]重新换回到原位置
        if(i > begin)
            swap(&nums[begin], &nums[i]);
    }
}
int** permute(int* nums, int numsSize, int* returnSize) {
    *returnSize=1;
    // size用来暂存numsSize，便于计算阶乘得到returnSize
    int size=numsSize;
    while(size>0){
        *returnSize = *returnSize * size;
        size--;
    }
    // 利用得到的returnSize和malloc即可对permutedNums进行内存分配
    int **permutedNums = (int**)malloc(*returnSize * sizeof(int*));
    // 对每个permutedNums[i]进行内存分配，用来存储每个全排列
    for(int i=0;i<*returnSize;i++)
        permutedNums[i] = (int*)malloc(numsSize * sizeof(int));
    // 再将returnSize归零，在handlePermute函数中需要该变量进行全排列计数
    *returnSize=0;
    handlePermute(nums, 0, numsSize, permutedNums, returnSize);
    return permutedNums;
}
```

## 遇到的问题及解决办法

在对permutedNums指针进行一维内存分配时，遇到错误：

```c
store to misaligned address 0x000200000001 for type 'int', which requires 4 byte alignment
```

原因是我在分配过程中，写为了：

```c
int **permutedNums = (int**)malloc(*returnSize * sizeof(int));
```

其中，malloc参数大小有误，由于permutedNums为指针的指针，所以其分配应为`(int**)malloc(*returnSize * sizeof(int*));`所以造成了该错误，引起内存泄漏。

