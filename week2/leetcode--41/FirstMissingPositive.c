void swap(int *a, int *b){
    int temp = *a;
    *a = *b;
    *b = temp;
}
int firstMissingPositive(int* nums, int numsSize) {
    for(int i=0;i<numsSize;i++){
        if(nums[i]<0 || nums[i]>numsSize)
            continue;
        else if(nums[nums[i]-1] != nums[i]){
            swap(&nums[nums[i]-1], &nums[i]);
            i--;
        }
    }
    for(int i=0;i<numsSize;i++)
        if(nums[i] != i+1)
            return i+1;
    return numsSize+1;
}