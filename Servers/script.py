import os
import sys
from multiprocessing import Process


def runServer(port):
    cmd = "node index.js " + port
    os.system(cmd)


if __name__ == "__main__":
    processes = list()
    num = int(sys.argv[1])

    for i in range(num):
        process = Process(target=runServer, args=(str(3000 + i), ))
        processes.append(process)
        process.start()

    for each in processes:
        each.join()