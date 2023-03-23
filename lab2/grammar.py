from copy import deepcopy
from dataclasses import dataclass


@dataclass
class Symbol:
    name: str
    spell: str = ""
    type: str = ""

    def __post_init__(self):
        self.spell = self.spell or self.name

    def __eq__(self, __o: "Symbol") -> bool:
        return self.name == __o.name

    def __str__(self) -> str:
        return self.spell

    def __repr__(self) -> str:
        return self.spell


class TermSymbol(Symbol):
    def __post_init__(self):
        self.type = "term"
        super().__post_init__()


class NonTermSymbol(Symbol):
    def __post_init__(self):
        self.type = "nonterm"
        super().__post_init__()


@dataclass
class TerminalSymbol(Symbol):
    pass


@dataclass
class Production:
    lhs: Symbol | str
    rhs: list[Symbol] | list[str]


@dataclass
class Grammar:
    term_symbols: list[TermSymbol]
    non_term_symbols: list[NonTermSymbol]
    productions: list[Production]
    start_symbol: Symbol

    def __post_init__(self):
        self.sync_productions()

    def symbol_production(self, s: Symbol):
        return [p for p in self.productions if p.lhs == s]

    def remove_symbol_productions(self, s: Symbol):
        self.productions = [p for p in self.productions if p.lhs != s]

    def add_production(self, lhs: Symbol, rhs: list[Symbol]):
        self.productions.append(Production(lhs, rhs))

    def symbol_names_dist(self):
        symbols = {}
        for s in self.term_symbols:
            symbols[s.name] = s
        for s in self.non_term_symbols:
            symbols[s.name] = s
        return symbols

    def sync_productions(self):
        symbols = self.symbol_names_dist()

        for prod in self.productions:
            if isinstance(prod.lhs, str):
                prod.lhs = symbols[prod.lhs]
            else:
                prod.lhs = symbols[prod.lhs.name]

            if isinstance(prod.rhs[0], str):
                prod.rhs = [symbols[s] for s in prod.rhs]
            else:
                prod.rhs = [symbols[s.name] for s in prod.rhs]

    def print(self):
        for s in self.non_term_symbols:
            productions = self.symbol_production(s)
            if not productions:
                continue

            print(s, "→", end=" ")
            str_prods = ["".join(rs.spell for rs in p.rhs) for p in productions]
            print("|".join(str_prods))

    def replace_previous_with_rhs(self, s: NonTermSymbol):
        idx = self.non_term_symbols.index(s)
        new_prods = []
        for p in self.symbol_production(s):
            first_rhs = p.rhs[0]
            if (
                first_rhs.type != "nonterm"
                or self.non_term_symbols.index(first_rhs) >= idx
            ):
                new_prods.append(p)
                continue
            for ref_p in self.symbol_production(first_rhs):
                new_prods.append(Production(p.lhs, ref_p.rhs + p.rhs[1:]))

        self.remove_symbol_productions(s)
        self.productions.extend(new_prods)


def reduce_left_recursion(old_g: Grammar) -> Grammar:
    g = deepcopy(old_g)
    for s in old_g.non_term_symbols:
        g.replace_previous_with_rhs(s)

        s_prods = g.symbol_production(s)
        alpha_rhs = [p.rhs[1:] for p in s_prods if p.rhs[0] == s]
        beta_rhs = [p.rhs for p in s_prods if p.rhs[0] != s]
        if not alpha_rhs:
            continue

        g.remove_symbol_productions(s)
        if not beta_rhs:
            for rhs in alpha_rhs:
                g.add_production(s, rhs + [s])
            continue

        alt_s = NonTermSymbol(s.name + "'")
        g.non_term_symbols.append(alt_s)
        for rhs in beta_rhs:
            g.add_production(s, rhs + [alt_s])
        for rhs in alpha_rhs:
            g.add_production(alt_s, rhs + [alt_s])
    return g
