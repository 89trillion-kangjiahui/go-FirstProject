import time
from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    wait_time = between(1, 2.5)

    @task
    def hello_world(self):
        self.client.get("/admin/select?code=xJu8uro9")
        self.client.get("/user/check?uid=99b4a72f645b49ef72f49b414e62f2c4&code=xJu8uro9")
        self.client.get("/user/register?username=赵阳")
        self.client.get("/user/login?uid=99b4a72f645b49ef72f49b414e62f2c4")
    def on_start(self):
        header = { "Content-Type":"application/x-www-form-urlencoded"}
        payload={
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
