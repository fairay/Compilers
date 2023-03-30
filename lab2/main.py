from grammar import Grammar, reduce_left_recursion, chomsky_normal_form
from dacite import from_dict
import json


def main():
    filename = input("Input file name:") or "data3.json"
    with open(filename, "r") as f:
        data = json.load(f)
    g = from_dict(data_class=Grammar, data=data)

    g.print()

    print("--------- Left Recursion Elimination -----------")
    g = reduce_left_recursion(g)
    g.print()

    print("--------- Chomsky Normal Form -----------")
    g = chomsky_normal_form(g)
    g.print()

if __name__ == "__main__":
    main()
