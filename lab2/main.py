from grammar import Grammar, reduce_left_recursion
from dacite import from_dict
import json


def main():
    with open("data3.json", "r") as f:
        data = json.load(f)
    result = from_dict(data_class=Grammar, data=data["grammar"])

    print(result)
    result.print()

    print("--------- Left Recursion Elimination -----------")
    g = reduce_left_recursion(result)
    g.print()


if __name__ == "__main__":
    main()
