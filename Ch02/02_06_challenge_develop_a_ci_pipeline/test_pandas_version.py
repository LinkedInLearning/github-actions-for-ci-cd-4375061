'''
Test for correct pandas version
'''
import pandas as pd


def test_pandas_version():
    ''' Use an assertion to check the output of pd.__version__ '''
    assert pd.__version__ in ["1.5.3"]


if __name__ == "__main__":
    if test_pandas_version() is True:
        print("Pandas version is correct!")
