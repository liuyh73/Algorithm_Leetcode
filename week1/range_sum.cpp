class Solution {
public:
    int countRangeSum(vector<int>& nums, int lower, int upper) {
        int sums[upper-lower+100]={0};
        int i=0;
        handleRange(nums, 0, nums.size(), lower, upper, i, sums);
    }
    void handleRange(vector<int>& nums, int left, int right, int &lower, int &upper, int &i, int sums[]){
        int sum=0;
        for(int j=left; j<=right; j++)
            sum+=nums[j];
        if(sum>=lower && sum<=upper)
            sums[i++]=sum
        if (left==right)    return;
        
    }
};