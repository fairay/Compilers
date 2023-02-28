import pytest

from dfa import build_dfa


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
    ]
)
def test_validate(message, is_valid):
    exp = "(a|b)*abb"
    dfa = build_dfa(exp)
    assert dfa.simulate(message) == is_valid
