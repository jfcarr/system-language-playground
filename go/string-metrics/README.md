# String Metrics

## Jaro-Winkler

In computer science and statistics, the Jaro–Winkler distance is a string metric measuring an edit distance between two sequences. It is a variant proposed in 1990 by William E. Winkler of the Jaro distance metric (1989, Matthew A. Jaro).

The Jaro–Winkler distance uses a prefix scale _p_ which gives more favourable ratings to strings that match from the beginning for a set prefix length _ℓ_.

The lower the Jaro–Winkler distance for two strings is, the more similar the strings are. The score is normalized such that 1 equates to no similarity and 0 is an exact match. The Jaro–Winkler similarity is given by 1 − Jaro–Winkler distance.  (The smetrics library flips this: 1 is an exact match, 0 is no similarity)

Although often referred to as a distance metric, the Jaro–Winkler distance is not a metric in the mathematical sense of that term because it does not obey the triangle inequality.

[Wikipedia article](https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance)

## Wagner-Fischer

In computer science, the Wagner–Fischer algorithm is a dynamic programming algorithm that computes the edit distance between two strings of characters.

The Wagner–Fischer algorithm computes edit distance based on the observation that if we reserve a matrix to hold the edit distances between all prefixes of the first string and all prefixes of the second, then we can compute the values in the matrix by flood filling the matrix, and thus find the distance between the two full strings as the last value computed.

[Wikipedia article](https://en.wikipedia.org/wiki/Wagner%E2%80%93Fischer_algorithm)
