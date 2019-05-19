import requests
import _thread
import json

def parse(text):
    return json.loads(text)

def get_session():
    r = requests.post(("http://localhost:8080/api/v1"), json={"endpoint": "identify", "method":"new"}, 
     headers={'Content-type': 'application/json'})
    print("""Response {} -- {} """.format(r.status_code, r.text))
    
    sessionKey = parse(r.text)["message"]
    print ("Aquired session key: " + sessionKey)
    r.close()
    return sessionKey

def test_endpoint(sessionKey):
    r = requests.post(("http://localhost:8080/api/v1"), json={"session": sessionKey, "endpoint": "cheese", "method":"hello", "payload": {"type":""}}, 
     headers={'Content-type': 'application/json'})
    print("""Response {} -- {} """.format(r.status_code, r.text))
    r.close()

key = get_session()

test_endpoint(key)
