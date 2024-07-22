# What is it?
program to break a sha256 hash that consists of multiple words in the Polish language for a challange from @sekurak on X/Twitter.

# How is it?
Multi threaded, written in GO.

# Issues?
- Null

# What are those files? 
- odm.txt -> Polish words grouped by the 'main' word others derive from.
- slowa.txt -> Every polish word.
- valid  -> the hash given by @sekurak consisting of 5 lowercase polish words with no spaces in between.

# Plans for the future?
- Split the work among multiple machines
- Have a "dashboard" that shows the progress on each machine
- scrape Polish internet and compile a dictionary of words based on their frequency in actual spoken Polish.
