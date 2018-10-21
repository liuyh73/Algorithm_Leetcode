# leetcode--50.Pow(x,n)题解

## 题目

Implement [pow(*x*, *n*)](http://www.cplusplus.com/reference/valarray/pow/), which calculates *x* raised to the power *n* ($x^n$).

**Example 1:**

```
Input: 2.00000, 10
Output: 1024.00000
```

**Example 2:**

```
Input: 2.10000, 3
Output: 9.26100
```

**Example 3:**

```
Input: 2.00000, -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
```

## 思路与解法

在该题目中，数据范围为-100.0<x<100.0，$-2^{31}$<=n<=$2^{31}$-1，所以普通的循环求幂的方法效率就会很低，所以可以采用快速求幂的方法。求数字x的n次幂：设中间结果为r，最终结果为ans：

第一遍计算：得到r=pow(x,2)，若n为偶数，则ans=$pow(x,2)^{n/2}$；若n为奇数，则ans=x*$pow(x,2)^{n/2}$；

第二遍计算：得到r=pow(r,2)，

​	若n为偶数：

​		若n/2为偶数，则ans=$pow(r,2)^{n/4}$，

​		若n/2为奇数，则ans=r * $pow(r,2)^{n/4}$；

​	若n为奇数：

​		若n/2为偶数，则ans=x * $pow(r,2)^{n/4}$，

​		若n/2为奇数，则ans=x *  r * $pow(r,2)^{n/4}$；

……

第$log_2n$遍计算：得到最终结果

以x=2，n=7为例：

第一遍计算：得到r=pow(2,2)，由于n为奇数，此时r=2*$pow(2,2)^{7/2}$，即为r=2 *${(2^2)}^3$；

第二遍计算：得到r=pow($2^2$,2)，由于n/2也为奇数，此时r=2 * $2^2$*$pow(2^2,2)^{3/2}$，即为r=2 * $2^2$ * $((2^2)^2)^1$；

化简即得到最终结果。

由上述分析可知，可以采用递归算法不断将n除以2，然后判断奇偶性来得到最终结果，时间复杂度为O(logn)。

## 代码实现

题目中给定了确实的函数：

```C
double myPow(double x, int n) {

}
```

实现如下：

```C
double myPow(double x, int n) {
    if(n==0)    return 1;
    if(n<0)     return (1/x)*myPow(1/x,-n-1);
    return n%2==0 ? myPow(x*x, n/2) : x*myPow(x*x, (n-1)/2);
}
```

上述函数中，

- if(n==0)表示数字x得0次幂，即为1；
- if(n<0)表示负数次幂，此时可以去1/x的-n次幂来计算；
- if(n>0)为一般计算，根据n的奇偶行来决定返回结果`myPow(x*x, n/2)` 或者`x*myPow(x*x, (n-1)/2);`。

## 遇到的问题及解决方法

在解题过程中，n<0情况刚开始的写法为：

```C
if(n<0)     return myPow(1/x,-n);
```

这种写法会出现溢出错误：原因是当输入数据为n=-2147483648时，-n超过了int表示范围。可以返回`(1/x)*myPow(1/x,-n-1);`来避免溢出。

这道题的难点在于快速幂的理解与递归实现、int溢出问题。