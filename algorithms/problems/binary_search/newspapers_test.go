package binarysearch

/*
Newspapers
https://algo.monster/problems/newspapers_split

You've begun your new job to organize newspapers.
Each morning, you are to separate the newspapers into smaller piles and
assign each pile to a co-worker. This way, your co-workers can read through
the newspapers and examine their contents simultaneously.

Each newspaper is marked with a read time to finish all its contents.
A worker can read one newspaper at a time, and, when they finish one,
they can start reading the next. Your goal is to minimize the amount of time
needed for your co-workers to finish all newspapers. Additionally, the newspapers
came in a particular order, and you must not disarrange the original ordering when
distributing the newspapers. In other words, you cannot pick and choose newspapers
randomly from the whole pile to assign to a co-worker, but rather you must take a
subsection of consecutive newspapers from the whole pile.

What is the minimum amount of time it would take to have your coworkers go through
all the newspapers? That is, if you optimize the distribution of newspapers, what is
the longest reading time among all piles?

Constraints
- 1 <= newspapers_read_times.length <= 10^5
- 1 <= newspapers_read_times[i] <= 10^5
- 1 <= num_coworkers <= 10^5

Examples
	Example 1:
	Input: newspapers_read_times = [7,2,5,10,8], num_coworkers = 2
	Output: 18
Explanation:
	Assign first 3 newspapers to one coworker then assign the rest to another.
	The time it takes for the first 3 newspapers is 7 + 2 + 5 = 14 and for
	the last 2 is 10 + 8 = 18.

Example 2:
	Input: newspapers_read_times = [2,3,5,7], num_coworkers = 3
	Output: 7
Explanation:
	Assign [2, 3], [5], and [7] separately to workers. The minimum time is 7.

==============================================================================

Mmk, gosh, where to begin? So, we're trying to split the newspapers into
separate piles such that it minimizes the work for all the readers.
We can't sort the input collection according to the problem.

So, given 7 2 5 10 8, we split that into 7+2+5 and 10+8.
Using brute force, we'd probably need to loop many times to find the maximum
read time.

If we have one worker, then we would just sum up the entire
array and return that as the minimum read time.

Now, this is a binary search problem, so how would we "know" to bisect
the array into 7 2 5 and 10 8?

It's like we should sum up the subarrays to a certain point, but how?
I guess we could do a prefix sum to make it the array 7 9 14 24 32.
This is like 7 9 14 and 10 18.
*/
