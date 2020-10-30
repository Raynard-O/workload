
# app.py - a minimal flask api using flask_restful
from flask import Flask, render_template, request, url_for, json
from flask_restful import Resource, Api
import person_pb2 as person
import requests
import uuid
import os
port = int(os.environ.get("PORT", 5000))
app = Flask(__name__)
api = Api(app)

RFWID = ""


@app.route('/')
def index():
    # Set up unique identifier ID for each request:
    RFWID = str(uuid.uuid4())

    return render_template("index.html", rfwid=RFWID)


@app.route('/response', methods=["POST"])
def response():
    # get data from form and set up parameters for post request
    url = "http://workload-env.eba-thtrpcmw.ca-central-1.elasticbeanstalk.com/client"
    params = {'RFWID': str(request.form.get("RFWID_form")),
              'BenchmarkType': str(request.form.get("BenchmarkType_form")),
              'WorkloadMetric': str(request.form.get("WorkloadMetric_form")),
              'BatchUnit': str(request.form.get("BatchUnit_form")),  # records in each batch
              'BatchID': str(request.form.get("BatchID_form")),  # offset of this request
              'BatchSize': str(request.form.get("BatchSize_form")),  # number of batches requested
              'BinarySerialization': str(request.form.get("BinarySeri_form"))}

    # BinarySerialization will send "binary" and "None" values

    # Find what workload metric the client sent
    def typeofresponse(type):
        types = {
            "CPU": 0,
            "NETIN": 1,
            "NETOUT": 2,
            "MEMUTI": 3,
            "FinalTarget": 4
        }
        return types.get(type)

    workload = typeofresponse(str(request.form.get("WorkloadMetric_form")))

    # Send POST request to server and save response into variable
    r = requests.request('POST', url, data=params)

    # return screen includes: RFW ID, Last Batch ID, Samples requested

    i = int(request.form.get("BatchUnit_form")) * int(request.form.get("BatchID_form"))

    if str(request.form.get("BinarySeri_form")) == "binary":
        # process return using protocolbuf
        Data = person.RFD()
        Data.ParseFromString(r.content)

        # Preparing parameters to pass to front end
        ReturnRFWID = Data.RFWID
        LastBatchID = Data.last_batch_id

        samples = []
        for batch in Data.batches:
            batchID = batch.Batch_ID
            for sample in batch.samples:
                i += 1
                if workload == 0:
                    samples.append(dict(id=i, batchid=batchID, sample=sample.CPU_utilization))
                elif workload == 1:
                    samples.append(dict(id=i, batchid=batchID, sample=sample.NetworkIN))
                elif workload == 2:
                    samples.append(dict(id=i, batchid=batchID, sample=sample.NetworkOUT))
                elif workload == 3:
                    samples.append(dict(id=i, batchid=batchID, sample=sample.Memory_utilization))
                elif workload == 4:
                    samples.append(dict(id=i, batchid=batchID, sample=sample.FinalTarget))
        Data = str(Data)
    else:
        # Deserializing the response from JSON to Python type dict
        Data = r.json()
        Data = Data["data"]
        # Preparing parameters to pass to front end
        ReturnRFWID = Data["RFWID"]
        LastBatchID = Data["last_batch_id"]

        samples = []
        for batch in Data["batches"]:
            batchID = batch["Batch_ID"]
            for sample in batch["samples"]:
                i += 1
                if workload == 0:
                    samples.append(dict(id=i, batchid=batchID, sample=sample["CPU_utilization"]))
                elif workload == 1:
                    samples.append(dict(id=i, batchid=batchID, sample=sample["NetworkIN"]))
                elif workload == 2:
                    samples.append(dict(id=i, batchid=batchID, sample=sample["NetworkOUT"]))
                elif workload == 3:
                    samples.append(dict(id=i, batchid=batchID, sample=sample["Memory_utilization"]))
                elif workload == 4:
                    samples.append(dict(id=i, batchid=batchID, sample=sample["FinalTarget"]))

        Data = str(Data)
    return render_template("response.html", ReturnRFWID=ReturnRFWID, LastBatchID=LastBatchID, samples=samples,
                           params=params)


class HelloWorld(Resource):
    def get(self):
        return {'hello': 'world'}


api.add_resource(HelloWorld, '/hello')

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=port)
