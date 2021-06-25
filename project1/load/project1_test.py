import time
from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    wait_time = between(1, 2.5)

    @task
    def hello_world(self):
        self.client.get("/soldier/getAll?rarity=1&unlockArena=0&cvc=1000")
        self.client.get("/soldier/getRarity?id=10101")
        self.client.get("/soldier/getAtk?id=10101")
        self.client.get("/soldier/getByCvc?cvc=1000")
        self.client.get("/soldier/getByUnlockArena")
