# Bob


def isEmpty(what):
    '''
    An Empty string is considered any string that consists of only whitespace.
    '''
    if what == '':
        return True
    return False


def isQuestion(what):
    '''
    A question is any string that ends in a ? Unless that string is also in ALL
    CAPS, in which case it is shouting.  Don't hanlde that case here since it's
    easy to have the main hey() function define the order of precedence.
    '''
    if what[-1] == "?":
        return True
    return False


def isShouting(what):
    '''
    Shouting is any string that is in ALL CAPS
    '''
    if what.isupper():
        return True
    return False


def hey(what):
    '''
    Respond to various inputs
    '''
    # Remove any extraneous white space from the ends of the string
    what = what.strip()

    if isEmpty(what):
        return "Fine. Be that way!"
    # shouting takes  precedence over questions so this case needs to be
    # handled first. A string that ends in a ? is considered YELLING if it's in
    # ALL CAPS, and not considered a question.
    elif isShouting(what):
        return "Whoa, chill out!"
    elif isQuestion(what):
        return "Sure."
    # Any other input gets a Whatever.
    return "Whatever."