# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import msg_pb2 as msg__pb2


class MetricServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SendMetrics = channel.unary_unary(
                '/msg.MetricService/SendMetrics',
                request_serializer=msg__pb2.Metric.SerializeToString,
                response_deserializer=msg__pb2.MetricsReply.FromString,
                )
        self.UploadAsset = channel.stream_unary(
                '/msg.MetricService/UploadAsset',
                request_serializer=msg__pb2.AssetUpload.SerializeToString,
                response_deserializer=msg__pb2.AssetUploadReply.FromString,
                )


class MetricServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def SendMetrics(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UploadAsset(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MetricServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'SendMetrics': grpc.unary_unary_rpc_method_handler(
                    servicer.SendMetrics,
                    request_deserializer=msg__pb2.Metric.FromString,
                    response_serializer=msg__pb2.MetricsReply.SerializeToString,
            ),
            'UploadAsset': grpc.stream_unary_rpc_method_handler(
                    servicer.UploadAsset,
                    request_deserializer=msg__pb2.AssetUpload.FromString,
                    response_serializer=msg__pb2.AssetUploadReply.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'msg.MetricService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class MetricService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def SendMetrics(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/msg.MetricService/SendMetrics',
            msg__pb2.Metric.SerializeToString,
            msg__pb2.MetricsReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UploadAsset(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_unary(request_iterator, target, '/msg.MetricService/UploadAsset',
            msg__pb2.AssetUpload.SerializeToString,
            msg__pb2.AssetUploadReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
