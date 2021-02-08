# HackerRank
## Queen's Attack II's Solution Summary
[https://www.hackerrank.com/challenges/queens-attack-2](https://www.hackerrank.com/challenges/queens-attack-2)

### O(k) Approach:

```
n  : size of the board
r_q: row of the queen
c_q: column of the queen
r_o: row of the obstacle
c_o: column of the obstacle
```

Formula to calculate every direction of n x n chess board:
```
[|] Top          [r1]  -> r_q - 1
[|] Bottom       [r2]  -> n - r_q
[-] Left         [c1]  -> c_q - 1
[-] Right        [c1]  -> n - c_q
[/] Bottom Left  [d11] -> min(r_q, c_q) - 1
[/] Top Right    [d22] -> n - max(r_q, c_q)
[\] Top Left     [d21] -> n - max(r_q, n+1-c_q)
[\] Bottom Right [d12] -> n - max(n+1-r_q, c_q)

for diagonal [\], since the r and c will always at the opposite direction toward n,
we can reverse the smaller number (x) using: "n + 1 - x" so all numbers will point toward n,
hence the formula of "[/] Top Right" when all numbers point toward n will apply.
```

To determine obstacle that align with queen's direction:
```
horizontal align -> r_o == r_q
vertical align   -> c_o == c_q
diagonal / align -> r_o - r_q == c_o - c_q
diagonal \ align -> r_o - r_q == (c_o - c_q) * -1
```

The rest is just calculating the distance between the queen to every direction and sum it, only nearest obstacle that align with queen's attack direction matters to the calculation.