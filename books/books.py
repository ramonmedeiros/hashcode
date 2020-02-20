import os
import sys


def read_input(filename):
    with open(filename) as fd:
        library = {}
        books, libraries, days = [int(item) for item in fd.readline().split()]

        score = [int(item) for item in fd.readline().split()]

        # iterate over libraries
        for lib in range(libraries):

            # parse info and store
            bookNumber, signupTime, shipCapacity = [int(item) for item in fd.readline().split()]
            bookList = [int(item) for item in fd.readline().split()]
            library[lib] = {"bookNumber": bookNumber,
                            "signupTime": signupTime,
                            "shipCapacity": shipCapacity,
                            "books": bookList
                            }
    return books, library, days


print(read_input(sys.argv[1]))
