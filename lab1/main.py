from dfa import build_dfa


def main():
    print("Regular expression: ")
    exp = input()
    dfa = build_dfa(exp)

    print("Input string: ")
    message = input()
    is_valid = dfa.simulate(message)
    print(f'String valid: {is_valid}')


if __name__ == '__main__':
    main()
