"""Functions used in preparing Guido's gorgeous lasagna.

Learn about Guido, the creator of the Python language:
https://en.wikipedia.org/wiki/Guido_van_Rossum

This is a module docstring, used to describe the functionality
of a module and its functions and/or classes.
"""


EXPECTED_BAKE_TIME = 40
PREPARATION_TIME = 2

def bake_time_remaining(elapsed_bake_time):
    """Calculate the bake time remaining.
    :param elapsed_bake_time: int - baking time already elapsed.
    :return: int - remaining bake time (in minutes) derived from 'EXPECTED_BAKE_TIME'.
    """
    
    return EXPECTED_BAKE_TIME - elapsed_bake_time

def preparation_time_in_minutes(number_of_layers):
    """Calculate the preparation time in minutes.
    :param number_of_layers: int - number of layers in your lasagna.
    :return: int - preparation time (in minutes) derived from number of layers * PREPARATION_TIME.
    """
    return number_of_layers * PREPARATION_TIME


#TODO: define the 'elapsed_time_in_minutes()' function below.
# Remember to add a docstring (you can copy and then alter the one from bake_time_remaining.)
def elapsed_time_in_minutes(number_of_layers, elapsed_bake_time):
    """Calculate the time you've been cooking.
    :param number_of_layers: int - number of layers in lasagna.
    :param elaspsed_bake_time: int - time spent baking.
    :return: int - total minutes spent cooking
    """

    return preparation_time_in_minutes(number_of_layers) + elapsed_bake_time
    