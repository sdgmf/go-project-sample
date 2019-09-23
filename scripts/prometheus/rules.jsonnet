local app = std.extVar('app');
local email = "xxx@xxx";

local rules = [{
    alert: "ServiceDown["+app+"]",
    expr:"absent(up{app=\""+app+"\"}) == 1",
    "for": "10s",
    labels: {
        email: email,
    },
    annotations: {
        description:"{{ $labels.app}}  has been down for more than 10 seconds."
    }
},{
    alert: "GRPCServerErrorThrottlingHigh["+app+"]",
    expr:"sum(rate(grpc_server_handled_total{app=\""+app+"\",grpc_type=\"unary\",grpc_code!=\"OK\"}[1m])) by (instance) > 0",
    "for": "10s",
    labels: {
        email: email,
    },
    annotations: {
        description: "{{$labels.instance}} has error request for 10 senconds (current value: {{ $value }}s)"
    }
},{
    alert: "GRPServerCLatencyThrottlingHigh["+app+"]",
    expr:"histogram_quantile(0.99,sum(rate(grpc_server_handling_seconds_bucket{app=\""+app+"\",grpc_type=\"unary\"}[1m])) by (instance,le)) > 0.2",
    "for": "10s",
    labels: {
        email: email,
    },
    annotations: {
        description: "{{ $labels.instance }} has a tp99 request latency above 200ms (current value: {{ $value }}s)"
    }
},{
     alert: "GRPCClientErrorThrottlingHigh["+app+"]",
     expr:"sum(rate(grpc_client_handled_total{app=\""+app+"\",grpc_type=\"unary\",grpc_code!=\"OK\"}[1m])) by (instance) > 0",
    "for": "10s",
    labels: {
        email: email,
    },
    annotations: {
        description: "{{$labels.instance}} has error request for 10 senconds (current value: {{ $value }}s)"
    }
 },{
     alert: "GRPCClientLatencyThrottlingHigh["+app+"]",
     expr:"histogram_quantile(0.99,sum(rate(grpc_client_handling_seconds_bucket{app=\""+app+"\",grpc_type=\"unary\"}[1m])) by (instance,le)) > 0.2",
     "for": "10s",
     labels: {
         email: email,
     },
     annotations: {
        description: "{{ $labels.instance }} has a tp99 request latency above 200ms (current value: {{ $value }}s)"
     }
 },{
     alert: "HTTPErrorThrottlingHigh["+app+"]",
     expr:"sum(rate(http_server_requests_seconds_count{app=\""+app+"\",code!=\"200\"}[1m])) by (instance) > 0",
     "for": "10s",
     labels: {
         email: email,
     },
     annotations: {
        description: "{{$labels.instance}} has error request for 10 senconds (current value: {{ $value }}s)"
     }
 },{
     alert: "HTTPLatencyThrottlingHigh["+app+"]",
     expr:"histogram_quantile(0.99,sum(rate(http_server_requests_seconds_bucket{app=\""+app+"\"}[1m])) by (instance,le)) > 0.2",
     "for": "10s",
     labels: {
         email: email,
     },
     annotations: {
        description: "{{ $labels.instance }} has a tp99 request latency above 200ms (current value: {{ $value }}s)"
     }
 }];

{
    groups: [
        {
            name:app,
            rules:rules
        }
    ]
}

