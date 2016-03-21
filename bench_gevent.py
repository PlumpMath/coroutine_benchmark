#!/bin/env python


import gevent
import time
import sys

tasks = []
c_running = 0
concurrency = 100

if len(sys.argv)>1:
    concurrency = int(sys.argv[1])

print "start testing, concurrency " + str(concurrency) + ": "

def worker(i):
    global c_running
    c_running = c_running + 1;

    start = int(time.time())
    while True:
        now = int(time.time())
        if now-start>600:
            break;
        gevent.sleep(2)
    c_running = c_running - 1;

def info():
    while True:
        gevent.sleep(0.5)
        print "RUNNING " + str(c_running) 

def schedule(n):
    global tasks
    for i in range(n):
        gevent.sleep( 0.001 )
        tasks.append( gevent.spawn(worker, i) )

tasks.append( gevent.spawn(schedule, concurrency) )
tasks.append( gevent.spawn(info) )
gevent.joinall( tasks )
