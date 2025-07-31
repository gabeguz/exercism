'''
A simple program to cound the word frequency in a given string.
'''

from collections import Counter


def word_count(string):
    '''
    Return the frequency of words in input string.
    '''
    cnt = Counter()
    for word in string.split():
        cnt[word] += 1

    return cnt