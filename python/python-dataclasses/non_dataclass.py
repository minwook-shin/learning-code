class User:
    def __init__(self):
        self.name = ""
        self.agree = True
        self.tag = []

    def __getitem__(self, idx):
        return self.tag[idx]

    def add_tag(self, tag):
        self.tag.append(tag)
