
# class Solution:
#     def isHappy(self, n: int) -> bool:
#         return self.simulation(n, set())

#     def sum_of_square_digits(self, num: int) -> int:
#         # tim so hang don vi
#         if num == 0:
#             return 0
        
#         return (num%10)**2 + self.sum_of_square_digits(num//10)

#     def simulation(self, number, seen):
#         if number == 1:
#             return True
#         if number in seen:
#             return False
        
#         seen.add(number)
#         next_number = self.sum_of_square_digits(number)

#         return self.simulation(next_number, seen)


class Solution:
    def isHappy(self, n: int) -> bool:
        seen = set()

        while n not in seen:
            seen.add(n)
            n = self.sum_of_square_digits(n)
            if n==1:
                return True
        
        return False
    
    def sum_of_square_digits(self, n: int) -> int:
        sum = 0
        while n>0:
            sum += (n%10)**2
            n //= 10 
        
        return sum



