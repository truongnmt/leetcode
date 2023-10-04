# 2,2,5,6,7,9
#     i
# count < 0 (count==2 and cost[i+1]>cost[i]: count=0; continue)
# totalCost = 9 7 5 2

class Solution:
    def minimumCost(self, cost: List[int]) -> int:
        # attemp1
        # if len(cost)<3:
        #     return sum(cost)

        # cost.sort()

        # count = 0
        # totalCost = 0

        # for i in range(len(cost)-1, -1, -1):
        #     if count<2:
        #         totalCost += cost[i]
        #         count += 1
        #     elif count==2 and cost[i+1]>=cost[i]:
        #         count = 0
        
        # return totalCost

        #--------------------------------------

        # attemp2
        cost.sort(reverse=True)
        totalSum = 0
        for i, c in enumerate(cost):
            if i%3 != 2:
                totalSum += c

        return totalSum


        