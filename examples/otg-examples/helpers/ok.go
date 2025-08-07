package helpers

import (
	"time"

	"github.com/open-traffic-generator/snappi/gosnappi"
)

type ExpectedBgpMetrics struct {
	Advertised int32
	Received   int32
}

type ExpectedIsisMetrics struct {
	L1SessionsUp   int32
	L2SessionsUp   int32
	L1DatabaseSize int32
	L2DatabaseSize int32
}
type ExpectedPortMetrics struct {
	FramesRx int32
}

type ExpectedFlowMetrics struct {
	FramesRx     int64
	FramesRxRate float32
}

type ExpectedState struct {
	Port map[string]ExpectedPortMetrics
	Flow map[string]ExpectedFlowMetrics
	Bgp4 map[string]ExpectedBgpMetrics
	Bgp6 map[string]ExpectedBgpMetrics
	Isis map[string]ExpectedIsisMetrics
}

func NewExpectedState() ExpectedState {
	e := ExpectedState{
		Port: map[string]ExpectedPortMetrics{},
		Flow: map[string]ExpectedFlowMetrics{},
		Bgp4: map[string]ExpectedBgpMetrics{},
		Bgp6: map[string]ExpectedBgpMetrics{},
		Isis: map[string]ExpectedIsisMetrics{},
	}
	return e
}

func (client *ApiClient) FlowMetricsOk(expectedState ExpectedState) (bool, error) {
	dNames := []string{}
	for name := range expectedState.Flow {
		dNames = append(dNames, name)
	}

	fMetrics, err := client.GetFlowMetrics(dNames, nil)
	if err != nil {
		return false, err
	}

	PrintMetricsTable(&MetricsTableOpts{
		FlowMetrics: fMetrics,
	})

	expected := true
	for _, f := range fMetrics.Items() {
		expectedMetrics := expectedState.Flow[f.Name()]
		if f.FramesRx() != expectedMetrics.FramesRx || f.FramesRxRate() != expectedMetrics.FramesRxRate {
			expected = false
		}
	}

	return expected, nil
}

func (client *ApiClient) AllBgp4SessionUp(expectedState ExpectedState) (bool, error) {
	dNames := []string{}
	for name := range expectedState.Bgp4 {
		dNames = append(dNames, name)
	}

	dMetrics, err := client.GetBgpv4Metrics(dNames, nil)
	if err != nil {
		return false, err
	}

	PrintMetricsTable(&MetricsTableOpts{
		Bgpv4Metrics: dMetrics,
	})

	expected := true
	for _, d := range dMetrics.Items() {
		expectedMetrics := expectedState.Bgp4[d.Name()]
		if d.SessionState() != gosnappi.Bgpv4MetricSessionState.UP || d.RoutesAdvertised() != expectedMetrics.Advertised || d.RoutesReceived() != expectedMetrics.Received {
			expected = false
		}
	}

	return expected, nil
}

func (client *ApiClient) AllBgp6SessionUp(expectedState ExpectedState) (bool, error) {
	dNames := []string{}
	for name := range expectedState.Bgp6 {
		dNames = append(dNames, name)
	}

	dMetrics, err := client.GetBgpv6Metrics(dNames, nil)
	if err != nil {
		return false, err
	}

	PrintMetricsTable(&MetricsTableOpts{
		Bgpv6Metrics: dMetrics,
	})

	expected := true
	for _, d := range dMetrics.Items() {
		expectedMetrics := expectedState.Bgp6[d.Name()]
		if d.SessionState() != gosnappi.Bgpv6MetricSessionState.UP || d.RoutesAdvertised() != expectedMetrics.Advertised || d.RoutesReceived() != expectedMetrics.Received {
			expected = false
		}
	}

	return expected, nil
}

func (client *ApiClient) AllIsisSessionUp(expectedState ExpectedState) (bool, error) {
	dNames := []string{}
	for name := range expectedState.Isis {
		dNames = append(dNames, name)
	}
	dMetrics, err := client.GetIsisMetrics(dNames, nil)
	if err != nil {
		return false, err
	}
	PrintMetricsTable(&MetricsTableOpts{
		IsisMetrics: dMetrics,
	})
	expected := true
	for _, d := range dMetrics.Items() {
		expectedMetrics := expectedState.Isis[d.Name()]
		if d.L1SessionsUp() != expectedMetrics.L1SessionsUp || d.L1DatabaseSize() != expectedMetrics.L1DatabaseSize || d.L1SessionsUp() != expectedMetrics.L1SessionsUp || d.L2DatabaseSize() != expectedMetrics.L2DatabaseSize {
			expected = false
		}
	}

	// TODO: wait explicitly until telemetry API (for talking to DUT) is available
	if expected {
		time.Sleep(2 * time.Second)
	}
	return expected, nil
}
