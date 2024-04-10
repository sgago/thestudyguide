package leetcode

/*
	535. Encode and Decode TinyURL
	https://leetcode.com/problems/encode-and-decode-tinyurl/description/

	Bummer, this is asking us to generate a tiny URL chunk of code.
	Not really what I wanted, in terms of leetcode problems.
	- There's some base 51 or 52 hashing used in crypto that'll
	and eliminate annoying characters like 1 and l. Good for users that are typing.
	- We can use any number of hashing algorithms to generate it,
	ideally create a low collision chance. I don't have these memorized.
	- Need collision handling. At scale, it can happen. We can gen a number for the DB to use like ID or similar.
	This won't be the first or last problem to run into resolving hash collisions.
	- Do NOT use simple numbers, users could just scan URLs and find them all.
	- FYI, a URL of 10^4 seems overly long and not recommended. Like <= 2K seems safe,
	> 2K might not work on all browsers. We can design a service for this still,
	but super long URLs will not work on all browsers. "Why doesn't my 10KB URL work when I share with my friend? Is this tiny URL site broken?"
	- There's hashing library out there somewhere that can do this for us.
*/
