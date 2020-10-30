from flask import Flask, render_template, request, url_for, json



import requests, uuid

url = "http://workload-env.eba-thtrpcmw.ca-central-1.elasticbeanstalk.com/client"

params = {'RFWID': "testing_request",
            'BenchmarkType': "DVD",
            'WorkloadMetric': "NETIN",
            'BatchUnit': 4, 
            'BatchID': 3,  
            'BatchSize': 2, 
            'BinarySerialization': "json"} 

r = requests.request('POST',url, data = params)

# json: 
Data = r.json()

# protobuf:
import person_pb2 as person

Data = person.RFD()
Data.ParseFromString(r.content)