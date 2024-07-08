# What is it?
program to break a sha256 hash that consists of multiple words in the Polish language for a challange from @sekurak on X/Twitter.
# How is it?
Multi threaded, written in GO.
# Issues?
- crypto/sha256 lib calculates a different hash with the same input if the input is long enough than most other tools. I suspect I have to enable the 'Boring' mode in the library.

# What are those files? 
- odm.txt -> Polish words grouped by the 'main' word others derive from.
- slowa.txt -> Every polish word.
- valid  -> the hash given by @sekurak consisting of 5 lowercase polish words with no spaces in between.
