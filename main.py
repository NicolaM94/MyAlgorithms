import time

'''Solves ProjectEuler's problem 40 - Champernown's costant'''
'''The problem requires to find the digits [1,10,100,1000,10000,100000,1000000] of the constant.'''
'''The idea is adding staks of numbers of lenght 10**n to a string and collect the digits in the requeired positions.'''
'''Last time elapsed for the solution: 1.5497207641601562e-05'''

# Keep tracks of the last numeber added to the string solution
number = 1
# Collects all the digits appended 
out = ""
# Collects the digits of the solution
collector = []

# The highest power of 10 required as solution is 10**6
for n in range(0,6):
    # we start adding numbers starting from the current 'number' to the power limit;
    # we stop adding digits to out when we reach the power limit for the current iteration 
    while len(out) < 10**n:
        # start adding numbers
        toadd = str(number)
        for a in toadd:
            out += a
        # increase 'number by one;
        number += 1
    # we add the second right most digit, because we need to keep count of the previous digits already added in 'out'
    collector.append(int(out[10**n-1]))

prod = 1
for c in collector:
    prod *= c

start = time.time()
print(prod)
print(time.time()-start)