from tree import (
    Term,
    Literal,
    ConcatTerm,
    RepeatTerm,
    parse
)
from typing import (
    Iterable
)

DEBUG = False


def _print(*args, **kwargs):
    if DEBUG:
        print(*args, **kwargs)


class DFA:
    def __init__(self) -> None:
        self.init_state: int = 0

        # State-transition table
        self._stt: dict[int, dict[str, int]] = {}
        self._finite_states: set[int] = set()

    def add_transition(self, current_st: int, next_st: int, symbol: str):
        self._stt.setdefault(current_st, {})[symbol] = next_st

    def add_finite_states(self, states: Iterable[int]):
        self._finite_states.update(states)

    def simulate(self, input_str: str) -> bool:
        state = self.init_state
        while len(input_str):
            _print(f'│{state}─ "{input_str}" ', end='')
            transition_symbol = input_str[0]
            input_str = input_str[1:]

            state_transitions = self._stt.get(state, {})
            state = state_transitions.get(transition_symbol, None)
            if state is None:
                _print('[no such transition]')
                return False
        _print(f'│{state}─ "{input_str}" ', end='')
        return state in self._finite_states

    def _split(self, R: set[int], C: set[int], a: str):
        R1: set[int] = set()
        R2: set[int] = set()
        for state in R:
            dest = self._stt[state].get(a, None)
            if dest in C:
                R1.add(state)
            else:
                R2.add(state)
        return R1, R2

    def _minimalize(self):
        P: list[set[int]] = []
        S = []

        def add_group(group: set[int]):
            P.append(group)
            for c in self.alphabet:
                S.append((group, c))

        def remove_group(group: set[int]):
            nonlocal S
            P.remove(group)
            S = [s for s in S if s[0] != group]

        add_group(self._finite_states)
        add_group(self.states.difference(self._finite_states))

        while len(S):
            C, a = S.pop(0)
            for R in P:
                R1, R2 = self._split(R, C, a)
                if R1 and R2:
                    remove_group(R)
                    add_group(R1)
                    add_group(R2)
                    break
        return P

    def minimalize(self):
        state_groups = self._minimalize()
        old_new = {}
        for new_state, old_states in enumerate(state_groups):
            for old_state in old_states:
                old_new[old_state] = new_state

        min_dfa = DFA()
        min_dfa.init_state = next(i for i, v in enumerate(state_groups) if self.init_state in v)
        min_dfa._finite_states = [i for i, v in enumerate(state_groups) if v.issubset(self._finite_states)]
        for src, transitions in self._stt.items():
            for symbol, dest in transitions.items():
                min_dfa.add_transition(old_new[src], old_new[dest], symbol)
        return min_dfa

    @property
    def alphabet(self):
        alphabet: set[str] = set()
        for transitions in self._stt.values():
            alphabet.update(transitions.keys())
        return alphabet

    @property
    def states(self):
        return set(self._stt.keys())


def calc_followpos(root: Term):
    followpos = {}

    def iterate(node: Term):
        if isinstance(node, Literal):
            followpos.setdefault(node.index, set())
            return

        for leaf in node.leafs:
            iterate(leaf)
        if isinstance(node, ConcatTerm):
            c1, c2 = node.leafs[0], node.leafs[1]
            for i in c1.lastpos():
                followpos.setdefault(i, set()).update(c2.firstpos())
        elif isinstance(node, RepeatTerm):
            for i in node.lastpos():
                followpos.setdefault(i, set()).update(node.firstpos())

    iterate(root)
    return followpos


def literal_table(root: Term):
    table = {}

    def iterate(node: Term):
        if isinstance(node, Literal):
            table[node.index] = node.char
        else:
            for leaf in node.leafs:
                iterate(leaf)

    iterate(root)
    return table


def dfa_from_tree(root: Term):
    followpos = calc_followpos(root)
    literals = literal_table(root)

    dfa = DFA()
    init_pos = root.firstpos()
    known_states = [init_pos]
    unprocessed = [init_pos]

    while len(unprocessed):
        nodes = unprocessed.pop(0)
        src_state = known_states.index(nodes)

        transuctions = {}
        for node in nodes:
            char = literals[node]
            fp = followpos[node]
            transuctions.setdefault(char, set()).update(fp)

        for char, node_set in transuctions.items():
            if char == '#':
                dfa.add_finite_states([src_state])
                continue

            if node_set not in known_states:
                known_states.append(node_set)
                unprocessed.append(node_set)
            dest_state = known_states.index(node_set)
            dfa.add_transition(src_state, dest_state, char)

    return dfa


def build_dfa(exp: str):
    root = parse(exp)
    raw_dfa = dfa_from_tree(root)
    min_dfa = raw_dfa.minimalize()
    return min_dfa
