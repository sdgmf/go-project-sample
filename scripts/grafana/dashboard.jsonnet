local grafana = import 'grafonnet-lib/grafonnet/grafana.libsonnet';
local dashboard = grafana.dashboard;
local row = grafana.row;
local singlestat = grafana.singlestat;
local prometheus = grafana.prometheus;
local graphPanel = grafana.graphPanel;
local template = grafana.template;
local row = grafana.row;

local app = std.extVar('app');

local baseUp() = singlestat.new(
  'Number of instances',
  datasource='Prometheus',
  span=2,
  valueName='current',
  transparent=true,
).addTarget(
  prometheus.target(
    'sum(up{app="' + app + '"})', instant=true
  )
);

local baseGrpcQPS() = singlestat.new(
  'Number of grpc request  per seconds',
  datasource='Prometheus',
  span=2,
  valueName='current',
  transparent=true,
).addTarget(
  prometheus.target(
    'sum(rate(grpc_server_handled_total{app="' + app + '",grpc_type="unary"}[1m]))',
    instant=true
  )
);

local baseGrpcError() = singlestat.new(
  'Percentage of grpc error request',
  format='percent',
  datasource='Prometheus',
  span=2,
  valueName='current',
  transparent=true,
).addTarget(
  prometheus.target(
    'sum(rate(grpc_server_handled_total{app="' + app + '",grpc_type="unary",grpc_code!="OK"}[1m])) /sum(rate(grpc_server_started_total{app="' + app + '",grpc_type="unary"}[1m])) * 100.0',
    instant=true
  )
);

local baseHttpQPS() = singlestat.new(
  'Number of http request  per seconds',
  datasource='Prometheus',
  span=2,
  valueName='current',
  transparent=true,
).addTarget(
  prometheus.target(
    'sum(rate(http_server_requests_seconds_count{app="' + app + '"}[1m]))',
    instant=true
  )
);

local baseHttpError() = singlestat.new(
  'Percentage of http error request',
  datasource='Prometheus',
  format='percent',
  span=2,
  valueName='current',
  transparent=true,
).addTarget(
  prometheus.target(
    'sum(rate(http_server_requests_seconds_count{app="' + app + '",code!="200"}[1m])) /sum(rate(http_server_requests_seconds_count{app="' + app + '"}[1m])) * 100.0',
    instant=true
  )
);

local goState(metric, description=null, format='none') = graphPanel.new(
  metric,
  span=6,
  fill=0,
  min=0,
  legend_values=true,
  legend_min=false,
  legend_max=true,
  legend_current=true,
  legend_total=false,
  legend_avg=false,
  legend_alignAsTable=true,
  legend_rightSide=true,
  transparent=true,
  description=description,
).addTarget(
  prometheus.target(
    metric + '{app="' + app + '"}',
    datasource='Prometheus',
    legendFormat='{{instance}}'
  )
);


local grpcQPS(kind='server', groups=['grpc_code']) = graphPanel.new(
  //title='grpc_' + kind + '_qps_' + std.join(',', groups),
  title='Number of grpc ' + kind + ' request  per seconds group by (' + std.join(',', groups) + ')',
  description='Number of grpc ' + kind + ' request per seconds  group by (' + std.join(',', groups) + ')',
  legend_values=true,
  legend_max=true,
  legend_current=true,
  legend_alignAsTable=true,
  legend_rightSide=true,
  transparent=true,
  span=6,
  fill=0,
  min=0,
).addTarget(
  prometheus.target(
    'sum(rate(grpc_' + kind + '_handled_total{app="' + app + '",grpc_type="unary"}[1m])) by (' + std.join(',', groups) + ')',
    datasource='Prometheus',
    legendFormat='{{' + std.join('}}.{{', groups) + '}}'
  )
);


local grpcErrorPercentage(kind='server', groups=['instance']) = graphPanel.new(
  //title='grpc_' + kind + '_error_percentage_' + std.join(',', groups),
  title='Percentage of grpc ' + kind + ' error request group by (' + std.join(',', groups) + ')',
  description='Percentage of grpc ' + kind + ' error request group by (' + std.join(',', groups) + ')',
  format='percent',
  legend_values=true,
  legend_max=true,
  legend_current=true,
  legend_alignAsTable=true,
  legend_rightSide=true,
  transparent=true,
  span=6,
  fill=0,
  min=0,
).addTarget(
  prometheus.target(
    'sum(rate(grpc_' + kind + '_handled_total{app="' + app + '",grpc_type="unary",grpc_code!="OK"}[1m])) by (' + std.join(',', groups) + ')/sum(rate(grpc_' + kind + '_started_total{app="' + app + '",grpc_type="unary"}[1m])) by (' + std.join(',', groups) + ')* 100.0',
    datasource='Prometheus',
    legendFormat='{{' + std.join('}}.{{', groups) + '}}'
  )
);


local grpcLatency(kind='server', groups=['instance'], quantile='0.99') = graphPanel.new(
  title='Latency of grpc ' + kind + ' request group by (' + std.join(',', groups) + ')',
  description='Latency of grpc ' + kind + ' request group by (' + std.join(',', groups) + ')',
  format='ms',
  legend_values=true,
  legend_max=true,
  legend_current=true,
  legend_alignAsTable=true,
  legend_rightSide=true,
  transparent=true,
  span=6,
  fill=0,
  min=0,
).addTarget(
  prometheus.target(
    '1000 * histogram_quantile(' + quantile + ',sum(rate(grpc_' + kind + '_handling_seconds_bucket{app="' + app + '",grpc_type="unary"}[1m])) by (' + std.join(',', groups) + ',le))',
    datasource='Prometheus',
    legendFormat='{{' + std.join('}}.{{', groups) + '}}'
  )
);


local httpQPS(kind='server', groups=['grpc_code']) = graphPanel.new(
  title='Number of http' + kind + ' request group by (' + std.join(',', groups) + ') per seconds',
  description='Number of http' + kind + ' request group by (' + std.join(',', groups) + ') per seconds',
  legend_values=true,
  legend_max=true,
  legend_current=true,
  legend_alignAsTable=true,
  legend_rightSide=true,
  transparent=true,
  span=6,
  fill=0,
  min=0,
).addTarget(
  prometheus.target(
    'sum(rate(http_server_requests_seconds_count{app="' + app + '"}[1m])) by (' + std.join(',', groups) + ')',
    datasource='Prometheus',
    legendFormat='{{' + std.join('}}.{{', groups) + '}}'
  )
);


local httpErrorPercentage(groups=['instance']) = graphPanel.new(
  //title='grpc_' + kind + '_error_percentage_' + std.join(',', groups),
  title='Percentage of http error request group by (' + std.join(',', groups) + ') ',
  description='Percentage of http error request group by (' + std.join(',', groups) + ')',
  format='percent',
  legend_values=true,
  legend_max=true,
  legend_current=true,
  legend_alignAsTable=true,
  legend_rightSide=true,
  transparent=true,
  span=6,
  fill=0,
  min=0,
).addTarget(
  prometheus.target(
    'sum(rate(http_server_requests_seconds_count{app="' + app + '",status!="200"}[1m])) by (' + std.join(',', groups) + ')/sum(rate(http_server_requests_seconds_count{app="' + app + '"}[1m])) by (' + std.join(',', groups) + ')* 100.0',
    datasource='Prometheus',
    legendFormat='{{' + std.join('}}.{{', groups) + '}}'
  )
);

local httpLatency(groups=['instance'], quantile='0.99') = graphPanel.new(
  title='Latency of http request group by (' + std.join(',', groups) + ')',
  description='Latency of http request group by (' + std.join(',', groups) + ')',
  format='ms',
  legend_values=true,
  legend_max=true,
  legend_current=true,
  legend_alignAsTable=true,
  legend_rightSide=true,
  transparent=true,
  span=6,
  fill=0,
  min=0,
).addTarget(
  prometheus.target(
    '1000 * histogram_quantile(' + quantile + ',sum(rate(http_server_requests_seconds_bucket{app="' + app + '"}[1m])) by (' + std.join(',', groups) + ',le))',
    datasource='Prometheus',
    legendFormat='{{' + std.join('}}.{{', groups) + '}}'
  )
);


dashboard.new(app, schemaVersion=16, tags=['go'], editable=true, uid=app)
.addPanel(row.new(title='Base', collapse=true)
          .addPanel(baseUp(), gridPos={ x: 0, y: 0, w: 4, h: 10 })
          .addPanel(baseGrpcQPS(), gridPos={ x: 4, y: 0, w: 4, h: 10 })
          .addPanel(baseGrpcError(), gridPos={ x: 8, y: 0, w: 4, h: 10 })
          .addPanel(baseHttpQPS(), gridPos={ x: 12, y: 0, w: 4, h: 10 })
          .addPanel(baseHttpError(), gridPos={ x: 16, y: 0, w: 4, h: 10 })
          , {})
.addPanel(row.new(title='Go', collapse=true)
          .addPanel(goState('go_goroutines', 'Number of goroutines that currently exist'), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_alloc_bytes', 'Number of bytes allocated and still in use'), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_alloc_bytes_total', 'Total number of bytes allocated, even if freed'), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_buck_hash_sys_bytes', 'Number of bytes used by the profiling bucket hash table'), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_frees_total', 'Total number of frees'), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_gc_cpu_fraction', "The fraction of this program's available CPU time used by the GC since the program started."), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_gc_sys_bytes', 'Number of bytes used for garbage collection system metadata'), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_heap_alloc_bytes', 'Number of heap bytes allocated and still in use'), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_heap_idle_bytes', 'Number of heap bytes waiting to be used'), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_heap_inuse_bytes', 'Number of heap bytes that are in use'), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_heap_objects', 'Number of allocated objects'), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_heap_released_bytes', 'Number of heap bytes released to OS'), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_heap_sys_bytes', 'Number of heap bytes obtained from system'), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_last_gc_time_seconds', 'Number of seconds since 1970 of last garbage collection'), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_lookups_total', 'Total number of pointer lookups'), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_mallocs_total', 'Total number of mallocs'), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_mcache_inuse_bytes', 'Number of bytes in use by mcache structures'), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_mcache_sys_bytes', 'Number of bytes used for mcache structures obtained from system'), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_mspan_inuse_bytes', 'Number of bytes in use by mspan structures'), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_mspan_sys_bytes', 'Number of bytes used for mspan structures obtained from system'), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_next_gc_bytes', 'Number of heap bytes when next garbage collection will take place'), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_other_sys_bytes', 'Number of bytes used for other system allocations'), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_stack_inuse_bytes', 'Number of bytes in use by the stack allocator'), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_stack_sys_bytes', 'Number of bytes obtained from system for stack allocator'), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(goState('go_memstats_sys_bytes', 'Number of bytes obtained from system'), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          , {})

.addPanel(row.new(title='Grpc Server request rate', collapse=true)
          .addPanel(grpcQPS('server', ['grpc_code']), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(grpcQPS('server', ['instance']), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(grpcQPS('server', ['grpc_service']), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(grpcQPS('server', ['grpc_service', 'grpc_method']), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          , {})
.addPanel(row.new(title='Grpc Server request error percentage', collapse=true)
          .addPanel(grpcErrorPercentage('server', ['grpc_service']), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(grpcErrorPercentage('server', ['grpc_service', 'grpc_method']), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(grpcErrorPercentage('server', ['instance']), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          , {})
.addPanel(row.new(title='Grpc server 99%-tile Latency of requests', collapse=true)
          .addPanel(grpcLatency('server', ['grpc_code'], 0.99), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(grpcLatency('server', ['instance'], 0.99), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(grpcLatency('server', ['grpc_service'], 0.99), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(grpcLatency('server', ['grpc_service', 'grpc_method'], 0.99), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          , {})
.addPanel(row.new(title='Grpc server 90%-tile Latency of requests', collapse=true)
          .addPanel(grpcLatency('server', ['grpc_code'], 0.90), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(grpcLatency('server', ['instance'], 0.90), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(grpcLatency('server', ['grpc_service'], 0.90), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(grpcLatency('server', ['grpc_service', 'grpc_method'], 0.99), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          , {})
.addPanel(row.new(title='Grpc client request rate', collapse=true)
          .addPanel(grpcQPS('client', ['grpc_code']), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(grpcQPS('client', ['instance']), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(grpcQPS('client', ['grpc_service']), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(grpcQPS('client', ['grpc_service', 'grpc_method']), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          , {})
.addPanel(row.new(title='Grpc client request error percentage', collapse=true)
          .addPanel(grpcErrorPercentage('client', ['grpc_service']), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(grpcErrorPercentage('client', ['grpc_service', 'grpc_method']), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(grpcErrorPercentage('client', ['instance']), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          , {})
.addPanel(row.new(title='Grpc client 99%-tile Latency of requests', collapse=true)
          .addPanel(grpcLatency('client', ['grpc_code'], 0.99), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(grpcLatency('client', ['instance'], 0.99), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(grpcLatency('client', ['grpc_service'], 0.99), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(grpcLatency('client', ['grpc_service', 'grpc_method'], 0.99), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          , {})
.addPanel(row.new(title='Grpc client 90%-tile Latency of requests', collapse=true)
          .addPanel(grpcLatency('client', ['grpc_code'], 0.90), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(grpcLatency('client', ['instance'], 0.90), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(grpcLatency('client', ['grpc_service'], 0.90), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(grpcLatency('client', ['grpc_service', 'grpc_method'], 0.99), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          , {})
.addPanel(row.new(title='Http server request rate', collapse=true)
          .addPanel(httpQPS(['grpc_code']), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(httpQPS(['instance']), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(httpQPS(['grpc_service']), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(httpQPS(['grpc_service', 'grpc_method']), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          , {})
.addPanel(row.new(title='Http server request error percentage', collapse=true)
          .addPanel(httpErrorPercentage(['instance']), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(httpErrorPercentage(['method', 'uri']), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          , {})
.addPanel(row.new(title='Http server 99%-tile Latency of requests', collapse=true)
          .addPanel(httpLatency(['status'], 0.99), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(httpLatency(['instance'], 0.99), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(httpLatency(['method', 'uri'], 0.99), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          , {})
.addPanel(row.new(title='Http server 90%-tile Latency of requests', collapse=true)
          .addPanel(httpLatency(['status'], 0.90), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          .addPanel(httpLatency(['instance'], 0.90), gridPos={ x: 12, y: 0, w: 12, h: 10 })
          .addPanel(httpLatency(['method', 'uri'], 0.90), gridPos={ x: 0, y: 0, w: 12, h: 10 })
          , {})
