
class SyntaxNode:
    pass

class Parser:
    def __init__(self):
        ind = 0

    def parse(self, s: str):
        if s[self.ind] == '(':
            with_bracket(s)


    def with_bracket(self, s: str):
        subtree = []

        ind += 1
        # subtree.append('(')

        subtree = self.parse_inside_brackets()

        ind += 1
        # subtree.append(')')
        return subtree




    def parse_inside_brackets(self):
        while s[self.ind] != ')':
            pass
        return None