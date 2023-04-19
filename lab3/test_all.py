import pytest
from main import validate


@pytest.mark.parametrize(
    'phrase, expected',
    [
        ("C", True),
        ("i", True),
        ("iC", False),
        ("(C)", True),
        ("C+", False),
        #
        ("notnotnotC", True),
        ("(not(not(i)))", True),
        #
        ("C*i/(notC)divC", True),
        ("C*i/(notCdivC", False),
        ("C*i/(notC)divC)", False),
        ("C*i/(modC)", False),
        #
        ("C/i<>C", True),
        ("C/i>C", True),
        ("(C/i>C)", False),
        ("C/i=C+(i*C*i)", True),
    ]
)
def test_parse(phrase, expected):
    assert validate(phrase) == expected


if __name__ == '__main__':
    pass
