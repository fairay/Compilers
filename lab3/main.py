class ParserError(RuntimeError):
    pass


class UnexpectedSymbolError(ParserError):
    def __init__(self, index: int) -> None:
        msg = f"unexpected symbol at position {index}"
        super().__init__(msg)


class NotYetError(ParserError):
    def __init__(self, s: str, index: int) -> None:
        msg = f"parsed expression with end at position {index} while length of the string is {len(s)}"
        super().__init__(msg)


class InternalParserError(ParserError):
    pass


class Parser:
    first = {
        "term": "iC(n",
        "_factor_brackets_expr": "(",
        "_factor_not": "n",
        "relationship_op": "=<>",
        "sign": "+-",
        "addition_op": "+-o",
        "multiplication_op": "*/dma",
        "identifier": "i",
        "constant": "C",
    }

    def __init__(self, s: str) -> None:
        self.i = 0
        self.s = s

    def in_first(self, func: str):
        return self.i < len(self.s) and self.s[self.i] in self.first[func]

    def is_next(self, symbol: str):
        return self.i < len(self.s) and self.s[self.i] == symbol

    def one_of_non_terms(self, *nt_list: str):
        for nt in nt_list:
            if self.s[self.i :].startswith(nt):
                self.i += len(nt)
                return
        raise InternalParserError()

    def run(self):
        try:
            self.S()
        except InternalParserError:
            raise UnexpectedSymbolError(self.i)

        if self.i != len(self.s):
            raise NotYetError(self.s, self.i)

    def S(self):
        self.expr()

    def expr(self):
        self.simple_expr()
        if self.in_first("relationship_op"):
            self.relationship_op()
            self.simple_expr()

    def simple_expr(self):
        if self.in_first("term"):
            self.term()
            self.simple_expr_()
        elif self.in_first("sign"):
            self.sign()
            self.term()
            self.simple_expr_()
        else:
            raise InternalParserError()

    def simple_expr_(self):
        if self.in_first("addition_op"):
            self.addition_op()
            self.term()
            self.simple_expr_()

    def term(self):
        self.factor()
        self.term_()

    def term_(self):
        if self.in_first("multiplication_op"):
            self.multiplication_op()
            self.factor()
            self.term_()

    def factor(self):
        if self.in_first("identifier"):
            self.identifier()
        elif self.in_first("constant"):
            self.constant()
        elif self.in_first("_factor_brackets_expr"):
            self._factor_brackets_expr()
        elif self.in_first("_factor_not"):
            self._factor_not()
        else:
            raise InternalParserError()

    def _factor_brackets_expr(self):
        self.one_of_non_terms("(")
        self.simple_expr()
        self.one_of_non_terms(")")

    def _factor_not(self):
        self.one_of_non_terms("not")
        self.factor()

    def relationship_op(self):
        self.one_of_non_terms("=", "<>", "<", "<=", ">", ">=")

    def sign(self):
        self.one_of_non_terms("+", "-")

    def addition_op(self):
        self.one_of_non_terms("+", "-", "or")

    def multiplication_op(self):
        self.one_of_non_terms("*", "/", "div", "mod", "and")

    def identifier(self):
        self.one_of_non_terms("i")

    def constant(self):
        self.one_of_non_terms("C")


class FullParser(Parser):
    def __init__(self, s: str) -> None:
        super().__init__(s)
        self.first.update({"tail": ";"})

    def S(self):
        self.program()

    def program(self):
        self.block()

    def block(self):
        self.one_of_non_terms("{")
        self.list_of_operators()
        self.one_of_non_terms("}")

    def list_of_operators(self):
        self.operator()
        self.tail()

    def tail(self):
        if self.in_first("tail"):
            self.one_of_non_terms(";")
            self.operator()
            self.tail()

    def operator(self):
        self.identifier()
        self.one_of_non_terms("=")
        self.expr()


def validate(s: str):
    p = Parser(s)
    try:
        p.run()
    except ParserError as e:
        print(f"parsing failed: {e}")
        return False
    return True


def full_validate(s: str):
    p = FullParser(s)
    try:
        p.run()
    except ParserError as e:
        print(f"parsing failed: {e}")
        return False
    return True


def main():
    pass


if __name__ == "__main__":
    main()
