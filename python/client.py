import json
import grpc
from msg_pb2_grpc import MetricServiceStub
from msg_pb2 import Metric, AssetUpload, AssetInfo
import time


channel = grpc.insecure_channel('localhost:9000')
stub = MetricServiceStub(channel)


metric_json = {
    "name": "experiment",
    "test": "indeed"
}
# stub.SendMetrics(Metric(
#     experimentID="dsads",
#     metricBody=json.dumps(metric_json),
#     createdAt=time.time_ns()
# ))


def read_bytes(file_, num_bytes):
    while True:
        bin = file_.read(num_bytes)
        if len(bin) != num_bytes:
            break
        yield bin


stub.UploadAsset(
    AssetUpload(info=AssetInfo(
        AssetName="test",
        AssetType=".py"
    ))
)

with open("msg_pb2.py   ", "rb") as f:
    for rec in read_bytes(f, 4):
        stub.UploadAsset(
            AssetUpload(
                content=rec
            )
        )
