int maxSubArray(int* nums, int numsSize){
    if(sizeof nums == 1) {
        return nums[0];
    }
    
    for(int i = 1; i < numsSize; i++) {
        if(nums[i-1] > 0) {
            nums[i] += nums[i-1];
        }
    }
    return getMax(nums, numsSize);
}

int getMax(int* nums, int numsSize) {
    int max = nums[0];
    for (int j = 1; j < numsSize; j++) {
        if (max < nums[j]) {
            max = nums[j];
        }
    }
    return max;
}