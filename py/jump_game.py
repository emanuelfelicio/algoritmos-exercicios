# o goal é ultimo index
# caso eu encontre um posição que chegue ou ultrapasse o goal, este será o meu novo goal
# pois se chego neste posso chegar no goal anterior
# se goal for == 0 ao final significa que posso chegar no ultimo index, pois este era o goal inicial
class Solution:
    def canJump(self, nums: List[int]) -> bool:
        goal = len(nums) - 1

        for i in range(len(nums) - 2, -1, -1):
            if i + nums[i] >= goal:
                goal = i
        return goal == 0
