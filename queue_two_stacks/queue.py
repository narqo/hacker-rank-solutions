class MyQueue(object):
    def __init__(self):
        self.first = []
        self.second = []

    def peek(self):
        self.prepare_second()
        return self.second[len(self.second)-1] if self.second else 0

    def pop(self):
        self.prepare_second()
        return self.second.pop() if self.second else 0

    def put(self, value):
        self.first.append(value)

    def prepare_second(self):
        if not self.second:
            while self.first:
                value = self.first.pop()
                self.second.append(value)

queue = MyQueue()
t = int(raw_input())
for line in xrange(t):
    values = map(int, raw_input().split())

    if values[0] == 1:
        queue.put(values[1])
    elif values[0] == 2:
        v = queue.pop()
    else:
        print queue.peek()

