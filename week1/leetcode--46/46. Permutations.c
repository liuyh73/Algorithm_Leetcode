/**
 * Return an array of arrays of size *returnSize.
 * Note: The returned array must be malloced, assume caller calls free().
 */
void swap(int *a, int *b){
    int temp = *a;
    *a = *b;
    *b= temp;
}
void handlePermute(int *nums, int begin, int numsSize, int *permutedNums[], int *returnSize){
    if(begin>=numsSize){
        for(int i=0;i<numsSize; i++) {
            permutedNums[*returnSize][i]=nums[i];
            //printf("%d", permutedNums[*returnSize][i]);
        }
        (*returnSize)++;
        //printf("--%d\n", *returnSize);
        return;
    }
    for(int i=begin; i<numsSize; i++) {
        if(i > begin)
            swap(&nums[begin], &nums[i]);
        handlePermute(nums, begin+1, numsSize, permutedNums, returnSize);
        if(i > begin)
            swap(&nums[begin], &nums[i]);
    }
}
int** permute(int* nums, int numsSize, int* returnSize) {
    *returnSize=1;
    int size=numsSize;
    while(size>0){
        *returnSize = *returnSize * size;
        size--;
    }
    // printf("%d\n", *returnSize);
    // int *permutedNums[*returnSize];
    int **permutedNums = (int**)malloc(*returnSize * sizeof(int*));
    for(int i=0;i<*returnSize;i++)
        permutedNums[i] = (int*)malloc(numsSize * sizeof(int));
    *returnSize=0;
    handlePermute(nums, 0, numsSize, permutedNums, returnSize);
    /*for(int j=0;j<*returnSize;j++){
        for(int i=0;i<numsSize; i++) {
            printf("%d", permutedNums[j][i]);
        }
        printf("\n");
    }*/
    return permutedNums;
}
