# low and high like a bucket in a hash map
# hash the ball number into bucket
# count the bucket size

class Solution:
    # def countBalls(self, lowLimit: int, highLimit: int) -> int:
    #     # create buckets
    #     boxes = defaultdict(list)
    #     maxBallCount = 0

    #     # hash the ball
    #     for i in range(lowLimit, highLimit+1):
    #         hashResult = self.hashBall(i)
    #         boxes[hashResult].append(i)
    #         maxBallCount = max(maxBallCount, len(boxes[hashResult]))
    #         # print(f"i={i} hashResult={hashResult} boxes[hashResult]={boxes[hashResult]} maxBallCount={maxBallCount}")

    #     return maxBallCount

    # def hashBall(self, num: int) -> int:
    #     sum = 0
    #     while num>0:
    #         sum += num%10
    #         num //= 10 # 321 / 10 = 32 rem 1; 32/10 = 3 rem 2; 3/10 = 0 rem 3

    #     return sum

    def countBalls(self, lowLimit: int, highLimit: int) -> int:
        freq = defaultdict(int)

        for num in range(lowLimit, highLimit+1):
            sumBall = sum(int(digit) for digit in str(num))
            freq[sumBall] += 1

        return max(freq.values())
