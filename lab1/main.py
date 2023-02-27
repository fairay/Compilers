from dfa import build_dfa, DFA
from tree import parse


def test_minimize():
    dfa = DFA()
    dfa.init_state = 0
    dfa.add_finite_states([2, 3, 4])
    dfa.add_transition(0, 1, '0')
    dfa.add_transition(0, 2, '1')
    dfa.add_transition(1, 0, '0')
    dfa.add_transition(1, 3, '1')
    dfa.add_transition(2, 4, '0')
    dfa.add_transition(2, 5, '1')
    dfa.add_transition(3, 4, '0')
    dfa.add_transition(3, 5, '1')
    dfa.add_transition(4, 4, '0')
    dfa.add_transition(4, 5, '1')
    dfa.add_transition(5, 5, '0')
    dfa.add_transition(5, 5, '1')

    min_dfa = dfa.minimalize()

def test():
    exp = "(a|b)*abb"
    dfa = build_dfa(exp)

    input_s = [
        "abababb",
        "abb",
        "aaabb",
        "bab",
        "babba"
        "",
        "bb"
    ]
    for s in input_s:
        print(s, dfa.simulate(s))


def main():
    return test()
    return test_minimize()
    dfa = DFA()
    dfa.init_state = 0
    dfa.add_finite_states([0])
    dfa.add_transition(0, 0, '0')
    dfa.add_transition(0, 1, '1')
    dfa.add_transition(1, 0, '1')
    dfa.add_transition(1, 2, '0')
    dfa.add_transition(2, 1, '0')
    dfa.add_transition(2, 2, '1')

    # print(dfa.simulate('1001'))
    # print(dfa.simulate('101'))
    # print(dfa.simulate(''))
    # print(dfa.simulate('00'))
    print(dfa.simulate('101011'))


if __name__ == '__main__':
    main()
