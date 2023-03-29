import dataclasses
import pytest
import json
from dacite import from_dict

from grammar import Grammar, reduce_left_recursion


def eq_grammar(a: Grammar, b: Grammar):
    assert a.start_symbol == b.start_symbol
    ap = set(p.to_str() for p in a.productions)
    bp = set(p.to_str() for p in b.productions)
    assert ap == bp
    assert set(a.term_symbols) == set(b.term_symbols)
    assert set(a.non_term_symbols) == set(b.non_term_symbols)


@pytest.mark.parametrize(
    'filename',
    [("1"), ("2")]
)
def test_left_recursion(filename):
    with open(f"test_data/src_{filename}.json", "r") as f:
        data = json.load(f)
        src_g = from_dict(data_class=Grammar, data=data)
    with open(f"test_data/expected_{filename}.json", "r") as f:
        data = json.load(f)
        expected_g = from_dict(data_class=Grammar, data=data)

    out_g = reduce_left_recursion(src_g)
    eq_grammar(out_g, expected_g)


@pytest.mark.parametrize(
    'filename',
    [("3")]
)
def test_chomsky_normal_form(filename):
    with open(f"test_data/src_{filename}.json", "r") as f:
        data = json.load(f)
        src_g = from_dict(data_class=Grammar, data=data)
    with open(f"test_data/expected_{filename}.json", "r") as f:
        data = json.load(f)
        expected_g = from_dict(data_class=Grammar, data=data)

    out_g = reduce_left_recursion(src_g)
    eq_grammar(out_g, expected_g)



if __name__ == '__main__':
    pass
