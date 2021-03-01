# HackerRank
## Number Line Jumps Solution Summary
[https://www.hackerrank.com/challenges/kangaroo](https://www.hackerrank.com/challenges/kangaroo)

Two kangooros has respectively position of x1 and x2 with velocity v1 and v2.

To determine if the two kangooros will able to meet at certain point of coordinate will can be determined by `(x2-x1) % (y2-y1) == 0`, if the difference is even, then they will be meet. 

But since at the challange stated that `the second kangaroo  has a starting location that is ahead the first kangooro and moves at a faster rate (meaning v2>v1) and is already ahead of the first kangaroo`. 

Then there is no way the first kangooro will catch up if `v1<v2`. so we can add "`&& v1>v2`" in the earlier formula.