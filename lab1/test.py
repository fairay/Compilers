import pytest

from dfa import build_dfa, DFA


@pytest.mark.parametrize(
    'message, is_valid',
    [
        ("abababb", True),
        ("abb", True),
        ("aaabb", True),
        ("bab", False),
        ("babba", False),
        ("", False),
        ("bb", False),
        ("ababbababb", True),
        ("ababbabab", False)
    ]
)
def test_validate(message, is_valid):
    exp = "(a|b)*abb"
    dfa = build_dfa(exp)
    assert dfa.simulate(message) == is_valid


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
    assert len(min_dfa.states) == 3


def test_minimize2():
    dfa = DFA()
    dfa.init_state = 0
    dfa.add_finite_states([4, 5])
    dfa.add_transition(0, 1, '0')
    dfa.add_transition(0, 2, '1')
    dfa.add_transition(1, 4, '0')
    dfa.add_transition(1, 5, '1')
    dfa.add_transition(2, 0, '0')
    dfa.add_transition(2, 0, '1')
    dfa.add_transition(3, 5, '0')
    dfa.add_transition(3, 4, '1')
    dfa.add_transition(4, 3, '0')
    dfa.add_transition(4, 5, '1')
    dfa.add_transition(5, 3, '0')
    dfa.add_transition(5, 4, '1')

    min_dfa = dfa.minimalize()
    assert len(min_dfa.states) == 4


@pytest.mark.parametrize(
    'message, is_valid',
    [
        ("1001", True),
        ("101", False),
        ("00", True),
        ("", True),
        ("bab", False),
        ("babba", False),
        ("bb", False),
    ]
)
def test_simulate(message, is_valid):
    dfa = DFA()
    dfa.init_state = 0
    dfa.add_finite_states([0])
    dfa.add_transition(0, 0, '0')
    dfa.add_transition(0, 1, '1')
    dfa.add_transition(1, 0, '1')
    dfa.add_transition(1, 2, '0')
    dfa.add_transition(2, 1, '0')
    dfa.add_transition(2, 2, '1')
    assert dfa.simulate(message) == is_valid


if __name__ == '__main__':
    pass
