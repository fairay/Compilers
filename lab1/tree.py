from enum import Enum


class Term:
    def __init__(self, *leafs: 'Term'):
        self.leafs = list(leafs)

    def nullable(self) -> bool:
        return True

    def firstpos(self) -> set[int]:
        return set()

    def lastpos(self) -> set[int]:
        return set()


class ConcatTerm(Term):
    def __init__(self, left: Term, right: Term):
        super().__init__(left, right)

    def __repr__(self):
        return f"{self.leafs[0]}.{self.leafs[1]}"

    def nullable(self) -> bool:
        return self.leafs[0].nullable() and self.leafs[1].nullable()

    def firstpos(self) -> set:
        s = self.leafs[0].firstpos()
        if self.leafs[0].nullable():
            s.update(self.leafs[1].firstpos())
        return s

    def lastpos(self) -> set:
        s = self.leafs[1].lastpos()
        if self.leafs[1].nullable():
            s.update(self.leafs[0].lastpos())
        return s


class ChoiceTerm(Term):
    def __init__(self, left: Term, right: Term):
        super().__init__(left, right)

    def __repr__(self):
        return f"({self.leafs[0]})|({self.leafs[1]})"

    def nullable(self) -> bool:
        return self.leafs[0].nullable() or self.leafs[1].nullable()

    def firstpos(self) -> set:
        s = self.leafs[0].firstpos()
        s.update(self.leafs[1].firstpos())
        return s

    def lastpos(self) -> set:
        s = self.leafs[1].lastpos()
        s.update(self.leafs[0].lastpos())
        return s


class RepeatTerm(Term):
    def __init__(self, left: Term):
        super().__init__(left)

    def __repr__(self):
        return f"({self.leafs[0]})*"

    def nullable(self) -> bool:
        return True

    def firstpos(self) -> set:
        return self.leafs[0].firstpos()

    def lastpos(self) -> set:
        return self.leafs[0].lastpos()


class Literal(Term):
    def __init__(self, char: str, index: int):
        super().__init__()
        self.char = char
        self.index = index + 1

    def __repr__(self):
        return f"{self.char}"

    def nullable(self) -> bool:
        return False

    def firstpos(self) -> set:
        return {self.index}

    def lastpos(self) -> set:
        return {self.index}


class EndLiteral(Literal):
    def __init__(self, index: int):
        super(EndLiteral, self).__init__('#', index)


class Operation(Enum):
    REPEAT = '*'
    CHOICE = '|'


class Parser:
    def __init__(self, s: str):
        self.ind = 0
        self.node_n = 0
        self.s = s

    def _parse_terms(self):
        assert self.cur == '('
        self.ind += 1

        terms = []
        while self.cur not in ')':
            if self.cur == '(':
                term = self.parse()
            elif self.cur == '*':
                term = Operation.REPEAT
            elif self.cur == '|':
                term = Operation.CHOICE
            else:
                term = Literal(self.cur, self.node_n)
                self.node_n += 1

            terms.append(term)
            self.ind += 1
        assert self.cur == ')'
        return terms

    def _repeat(self, terms):
        cterms = []
        for term in terms:
            if term == Operation.REPEAT:
                assert isinstance(cterms[-1], Term)
                cterms[-1] = RepeatTerm(cterms[-1])
            else:
                cterms.append(term)
        return cterms

    def _concat(self, terms):
        cterms = []
        for term in terms:
            if len(cterms) and isinstance(cterms[-1], Term) and isinstance(term, Term):
                cterms[-1] = ConcatTerm(cterms[-1], term)
            else:
                cterms.append(term)
        return cterms

    def _choice(self, terms):
        cterms = []
        upcoming_op = None
        for term in terms:
            if term == Operation.CHOICE:
                upcoming_op = term
            elif upcoming_op == Operation.CHOICE:
                cterms[-1] = ChoiceTerm(cterms[-1], term)
                upcoming_op = None
            else:
                cterms.append(term)
        return cterms

    def parse(self):
        terms = self._parse_terms()
        terms = self._repeat(terms)
        terms = self._concat(terms)
        terms = self._choice(terms)
        return terms[0]

    @property
    def cur(self):
        return self.s[self.ind]


def parse(s: str):
    s = f'({s})'
    parser = Parser(s)
    root = parser.parse()
    return ConcatTerm(root, EndLiteral(parser.node_n))


def main():
    exp = "(ab(cde)*(fg)hi)"
    exp = "(ab*c(ab)*)"
    exp = "(c(ab|c*a(b|c))*|bc*)"
    exp = "((a|b)*abb)"

    parser = Parser(exp)
    print(exp)
    tree = parser.parse()
    print(tree)
    print(tree.firstpos(), tree.lastpos())


if __name__ == '__main__':
    main()
