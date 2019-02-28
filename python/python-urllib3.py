import urllib3
from json import loads
from typing import Dict

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)


def getData(url: str) -> Dict:
    http = urllib3.PoolManager()
    req = http.request('GET', url)
    data = req.data
    if req.status == 200:
        source: Dict = loads(req.data.decode('utf-8'))
        return source
    else:
        return None


get = getData("http://localhost:3000/user/0")
print(get)
print(type(get) is dict)
