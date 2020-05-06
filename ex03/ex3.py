#This will display the prompt
print "I will now count my chickens:"

# Prints a text string, then performs calculations, divides 30 by 6 to equal 5 then adds 25
print "Hens", 25 + 30 / 6
# Prints a text string, then performs calculations, 25 multiplied by 3, then the remainer of 75 divided by 4 which equals 2 (the quotient is dismissed and the value shown is the remainder). This number is then subtracted from 100 to equal 98.
print "Roosters", 100 - 25 * 3 % 4

print "Now I will count the eggs:"

# This performs a calculation then displays result. Following PEMDAS, 4 % 2 is 0, then 1/4 is supposed to be .25 however due to floating point this becomes x, as I am unsure of how it is evaluated in the expression.
# Then proceding with the AS operations, 3 + 2 + 1 - 5 - x + 6 = 7 , x = 0. This means that the floating point did not include .25 as an integer. 
print 3 + 2 + 1 - 5 + 4 % 2 - 1 / 4 + 6

# Displays text string
print "Is it true that 3 + 2 < 5 - 7?"

# Displays the output of the calculations
print 3 + 2 < 5 - 7

# Displays text string, then the computational ouptut
print "What is 3 + 2?", 3 + 2
print "What is 5 - 7?", 5 - 7

# Displays text string
print "Oh that's why it's False."

# Displays text string
print "How about some more."

# Displays text string, then evaluates the expression as True or False
print "Is it greater?", 5 > -2
print "Is it greater or equal?",  5 >= -2
print "Is it less or equal?", 5 <= -2