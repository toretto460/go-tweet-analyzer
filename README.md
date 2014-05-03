go-tweet-analyzer
=================

experimental tweet analyzer

### The tf-idf strategy

#####the tool uses the tf-idf to detemine the word's weight

The ***tf*** is the ***term-frequency***, the number of term's occurrences within the tweet.
It is calculated in the simplest way (the number of the times that a term X occurs in a tweet T)

The ***idf*** is the inverse-document-frequency, it is used to adjust the weight for the very frequent terms like *the*, *of* ...

For example within the text "the cat is purple" the terms "the" and "is" are very commons so its frequency within a collection of documents will be very hight! 
However the most significant terms are "cat" and "purple". 
The ***idf*** is the statistic that shows how much information provides a term weight in a documents collection.
This value is obtained by dividing the total number of documents by the number of documents containing the term, and then taking the logarithm of that quotient.

idf(X) = log (total docs / docs with term X)