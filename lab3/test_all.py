import pytest
from main import validate, full_validate


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


@pytest.mark.parametrize(
    'phrase, expected',
    [
        ("{}", False),
        ("{i=C}", True),
        ("{i=C;}", False),
        ("{i=C;i=C<(i+i*Cor(notC))}", True),
        ("{i=i;i=C=i}", True),
    ]
)
def test_full_parse(phrase, expected):
    assert full_validate(phrase) == expected


if __name__ == '__main__':
    pass
