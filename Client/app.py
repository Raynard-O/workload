from flask import Flask, render_template, request, url_for, json

import person_pb2 as person

import requests, uuid

app = Flask(__name__)

from flask import json

# @app.route('/summary')def summary():
#     data = make_summary()
#     response = app.response_class(
#         response=json.dumps(data),
#         status=200,
#         mimetype='application/json'
#     )
#     return response



@app.route('/')
def index():
    return render_template("index.html")

@app.route('/response', methods=["POST"])
def response():
    # Set up unique identifier ID for each request:
    RFWID = str(uuid.uuid4())

    # get data from form and set up parameters for post request 
    url = "http://localhost:8080/client"
    params = {'RFWID': RFWID,
            'BenchmarkType': str(request.form.get("BenchmarkType_form")),
            'WorkloadMetric': str(request.form.get("WorkloadMetric_form")),
            'BatchUnit': str(request.form.get("BatchUnit_form")),
            'BatchID': str(request.form.get("BatchID_form")),
            'BatchSize': str(request.form.get("BatchSize_form")),
            'BinarySerialization': str(request.form.get("BinarySeri_form"))} 
            #BinarySerialization will send "binary" and "None" values

    #Find what workload metric the client sent
    def typeofresponse(type):
        types = {
            "CPU":0,
            "NetworkIn":1,
            "NetworkOut":2,
            "Memory":3,
            "FinalTarget":4
        }
        return types.get(type)
    
    workload = typeofresponse(str(request.form.get("WorkloadMetric_form")))

    #Send POST request to server and save response into variable
    r = requests.request('POST',url, data = params)

    #return screen includes: RFW ID, Last Batch ID, Samples requested

    if str(request.form.get("BinarySeri_form")) == "binary":
        #process return using protocolbuf
        Data = person.RFD()
        Data.ParseFromString(r.content)

        #Preparing parameters to pass to front end
        LastBatchID = Data.last_batch_id

        samples = []
        for batch in Data.batches:
            batchID = batch.Batch_ID
            for sample in batch.samples:
                if workload == 0:
                    samples.append([batchID,sample.CPU_utilization])
                elif workload == 1:
                    samples.append([batchID,sample.NetworkIN])
                elif workload == 2:
                    samples.append([batchID,sample.NetworkOUT])
                elif workload == 3:
                    samples.append([batchID,sample.Memory_utilization])
                elif workload == 4:
                    samples.append([batchID,sample.FinalTarget])

    else:
        #process return using JSON
        Data = r.json()
    
        #Preparing parameters to pass to front end
        
    
    return Data  #render_template("response.html",requestarray=requestarray)

from flask import json
@app.route('/summary')
def summary():
    r = requests.request('GET',"http://localhost:8080/summary")
    b = r.json()

    return str(b)



if __name__ == "__main__":
    app.run(debug=True)