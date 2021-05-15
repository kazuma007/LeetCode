// https://leetcode.com/submissions/detail/493678679/

void moveZeroes(int* nums, int numsSize){
    for (int i = 0; i < numsSize - 1; i++) {
      if (nums[i] == 0) {
        for (int j = i + 1; j < numsSize; j++) {
          if (nums[j] != 0) {
            int temp = nums[i];
            nums[i] = nums[j];
            nums[j] = temp;
            break;        
          }
        }
      }
    }
}