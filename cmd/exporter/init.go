package main

import (
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/nlconntrack"
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/nlqdisc"
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/procfd"
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/procio"
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/procipvs"
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/procnetdev"
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/procnetstat"
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/procsched"
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/procsnmp"
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/procsock"
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/procsoftnet"
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/proctcpsummary"
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/tracebiolatency"
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/tracekernel"
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/tracenetiftxlatency"
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/tracenetsoftirq"
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/tracepacketloss"
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/tracesocketlatency"
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/tracetcpreset"
	_ "github.com/alibaba/kubeskoop/pkg/exporter/probe/tracevirtcmdlat"
)
