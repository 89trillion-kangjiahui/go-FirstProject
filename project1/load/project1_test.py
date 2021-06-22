import time
from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    wait_time = between(1, 2.5)

    @task
    def hello_world(self):
        self.client.get("/soldier/getAll?rarity=1&unlockArena=0")
        self.client.get("/soldier/getRarity?id=10101")
        self.client.get("/soldier/atc?id=10101")
        self.client.get("/soldier/getAll/unlockArena")
