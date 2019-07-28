import requests
import _thread
import json

def parse(text):
    return json.loads(text)

def get_session():
    r = requests.post(("http://localhost:9489/api/v1"), json={"endpoint": "identify", "method":"new"}, 
     headers={'Content-type': 'application/json'})
    print("""Response {} -- {} """.format(r.status_code, r.text))
    
    sessionKey = parse(r.text)["message"]
    print ("Aquired session key: " + sessionKey)
    r.close()
    return sessionKey

def test_endpoint(sessionKey):
    r = requests.post(("http://localhost:9489/api/v1"), json={"session": sessionKey, "endpoint": "cheese", "method":"hello", "payload": {"type":""}}, 
     headers={'Content-type': 'application/json'})
    print("""Response {} -- {} """.format(r.status_code, r.text))
    r.close()


def test_endpoint_service_status():
    r = requests.post(("http://localhost:9489/api/v1"), json={"endpoint": "service", "method":"Status", "payload": {"id":"steve", "access_key":"123", "type":"alive"}}, 
     headers={'Content-type': 'application/json'})
    print("""Response {} -- {} """.format(r.status_code, r.text))
    r.close()

def test_endpoint_service_subscribe():
    r = requests.post(("http://localhost:9489/api/v1"), json={"endpoint": "service", "method":"Subscribe", "payload": {"login_key":"test", "to":"steve"}}, 
     headers={'Content-type': 'application/json'})
    print("""Response {} -- {} """.format(r.status_code, r.text))
    r.close()

def test_endpoint_service_getstatus():
    r = requests.post(("http://localhost:9489/api/v1"), json={"endpoint": "service", "method":"GetStatus", "payload": {"login_key":"test", "to":"steve"}}, 
     headers={'Content-type': 'application/json'})
    print("""Response {} -- {} """.format(r.status_code, r.text))
    r.close() 

test_endpoint_service_status()
test_endpoint_service_getstatus()

#key = get_session()

#test_endpoint(key)
