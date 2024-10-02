from typing import Deque


class RecentCounter:

    def __init__(self):
        self.queue = Deque()

    def ping(self, t: int) -> int:
        self.queue.append(t)

        window = t - 3000
        while self.queue and self.queue[0] < window:
            self.queue.popleft()

        return self.queue[0]
