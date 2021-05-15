// https://leetcode.com/problems/move-zeroes/submissions/

void moveZeroes(int* nums, int numsSize){
  int ans[numsSize];

  int ansNum = 0;
  for (int i = 0; i < numsSize; i++) {
    if (nums[i] != 0) {
      ans[ansNum] = nums[i];
      ansNum++;
    }
  }

  for (int j = ansNum; j < numsSize; j++) {
    ans[j] = 0;
  }

  for (int k = 0; k < numsSize; k++) {
    nums[k] = ans[k];
  }
}