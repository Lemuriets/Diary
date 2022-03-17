import requests

def test_auth(login: str, password: str) -> requests.Response:
    return requests.post("http://127.0.0.1:8000/auth/sign-up", data = {
        "login": login,
        "password": password,
    })


print(test_auth("testlogin", "testpass"))
