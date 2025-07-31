'''
A simple program to cound the word frequency in a given string.
'''


def word_count(string):
    '''
    Return the frequency of words in input string.
    '''
    wordlist = {}
    for word in string.split():
        if word not in wordlist:
            wordlist[word] = 1
        else:
            wordlist[word] = wordlist[word] + 1
    return wordlist