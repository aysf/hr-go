## Algorithm Skills

- sorted slices based on absolute value


## Solution

Discount Tags:
  First, it will always be optimal to choose as many positive elements as possible.
  Then, to make the sum even, there are two options, either remove an odd positive element or add a negative odd element.
  So, find the minimum amongst the absolute values of all odd elements, and update the answer accordingly.
