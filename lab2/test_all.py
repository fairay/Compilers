import dataclasses
import pytest
import json
from dacite import from_dict

from grammar import Grammar, reduce_left_recursion, chomsky_normal_form


def eq_grammar(a: Grammar, b: Grammar):
    assert a.start_symbol == b.start_symbol
    ap = set(p.to_str() for p in a.productions)
    bp = set(p.to_str() for p in b.productions)
    assert ap == bp
    assert set(a.term_symbols) == set(b.term_symbols)
    assert set(a.non_term_symbols) == set(b.non_term_symbols)


@pytest.mark.parametrize(
    'filename',
    ["left_recursion1", "left_recursion2"]
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
    ["chomsky1"]
)
def test_chomsky_normal_form(filename):
    with open(f"test_data/src_{filename}.json", "r") as f:
        data = json.load(f)
        src_g = from_dict(data_class=Grammar, data=data)
    with open(f"test_data/expected_{filename}.json", "r") as f:
        data = json.load(f)
        expected_g = from_dict(data_class=Grammar, data=data)

    out_g = chomsky_normal_form(src_g)
    eq_grammar(out_g, expected_g)


@pytest.mark.parametrize(
    'filename',
    ["full1"]
)
def test_full(filename):
    with open(f"test_data/src_{filename}.json", "r") as f:
        data = json.load(f)
        src_g = from_dict(data_class=Grammar, data=data)

    with open(f"test_data/expected_{filename}.json", "r") as f:
        data = json.load(f)
        expected_g = from_dict(data_class=Grammar, data=data)

    out_g = reduce_left_recursion(src_g)
    out_g = chomsky_normal_form(out_g)
    eq_grammar(out_g, expected_g)


if __name__ == '__main__':
    pass
