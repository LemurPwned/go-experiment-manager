import json
import grpc
from msg_pb2_grpc import MetricServiceStub
from msg_pb2 import Metric
import time


channel = grpc.insecure_channel('localhost:9000')
stub = MetricServiceStub(channel)


metric_json = {
    "name": "experiment",
    "test": "indeed"
}
stub.SendMetrics(Metric(
    experimentID="dsads",
    metricBody=json.dumps(metric_json),
    createdAt=time.time_ns()
))
