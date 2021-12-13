import sys
import random
import requests
from time import time
from concurrent.futures import ThreadPoolExecutor

PORT = "8080"
TYPE = "http"
HOST = "127.0.0.1"
LIM = 1000000


def randomRequestID():
    return random.randint(0, LIM)


def makeRequest(requestID):
    ID = str(requestID)
    URL = TYPE + "://" + HOST + ":" + PORT + "/" + ID
    response = requests.get(URL)
    print(response.text)


def timer(function, args):
    startTime = time()
    function(*args)
    endTime = time()
    totalTime = endTime - startTime
    print("%s seconds" % totalTime)


def main(num):
    with ThreadPoolExecutor(max_workers=100) as executor:
        tasks = [randomRequestID() for _ in range(num)]
        executor.map(makeRequest, tasks)
        executor.shutdown(wait=True)


if __name__ == "__main__":
    num = int(sys.argv[1])
    timer(main, [num])