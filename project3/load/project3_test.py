import time
from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    wait_time = between(1, 2.5)

    @task
    def hello_world(self):
        self.client.get("/admin/select?code=L4W5NGCd")
        self.client.get("/user/check?uid=3&code=L4W5NGCd")
        
    def on_start(self):
        header = { "Content-Type":"application/x-www-form-urlencoded"}
        payload={
            "uid":"1",
            "codeType":"2",
            "des": "牛逼",
            "receiveNum":"1000",
            "usefulDate": "2021-06-29 15:04:05",
            "jewel":"8",
            "gold": "4",
            "props": "5",
            "hero": "6",
            "batman": "8"
        }

        self.client.post("/admin/create",data=payload ,headers=header)
