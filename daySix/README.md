### Day 6

Super cool puzzle today and is an amazing task for less experienced developers as the quickest to implement solution is not the most efficient one.... by far

My first 'solution' was a complete waste of time. Instead of using the numbers `0` => `8` as keys I created an array which expands exponentially as the number of fish increases. 

Not only does this mean that it's going to use a metric ton of memory, but it also means that the solution is going to be very slow. As you have to after every day iterate through the array to adjust the values...

After treating the fish as a set of fish which correspond to the keys `0` through to `8`

We can simply shift down everything in the array one space then add the new fish to their respective position...

This solution improved efficiency by 99% and reduced the time to completion from ~1.5 mins to 307.375µs in part 1.

Then a simple adjustment to one number gave me puzzle two's solution