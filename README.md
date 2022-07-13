# Equation-Solver-With-Golang

Problem Description:

Write a golang program which could evaluate equations like below:

  => ((1+2)*(4/2))
  Ans: 6


Approch:
  1. Took input from user in infix order.
      Example: ((1+2)*(4/2))
  2. Convert Infix order To Postfix order.
      Example: ((1+2)*(4/2)) = 12+42/*
  3. Calculate Postfix result using Stack.
  
 
Reference & Resources:
  1. https://www.tutorialspoint.com/Convert-Infix-to-Postfix-Expression#:~:text=To%20convert%20infix%20expression%20to,maintaining%20the%20precedence%20of%20them.
  2. https://www.w3schools.com/go/
