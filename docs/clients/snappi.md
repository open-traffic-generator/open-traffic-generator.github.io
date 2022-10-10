# snappi

![snappi](https://github.com/open-traffic-generator/snappi/raw/main/snappi-logo.png)

## Overview

Test scripts written in [`snappi`](https://github.com/open-traffic-generator/snappi), an auto-generated Python module, can be executed against any traffic generator conforming to the Open Traffic Generator API. The examples of `snappi` coding patterns can also be found as part of [Ixia-c documentation](https://github.com/open-traffic-generator/ixia-c/blob/main/docs/snappi-constructs.md).

## Install on a client

```sh
python -m pip install --upgrade snappi
```

## Start scripting

Add a new file `hello.py` with following snippet:

```python
import snappi
# create a new API instance where location points to controller.
# this will use HTTP transport by default; in order to use gRPC instead,
# one can pass additional kwarg `transport=snappi.Transport.GRPC`
api = snappi.api(location='https://localhost')

# create a config object to be pushed to controller
config = api.config()
# add a port with location pointing to traffic engine
prt = config.ports.port(name='prt', location='localhost:5555')[-1]
# add a flow and assign endpoints
flw = config.flows.flow(name='flw')[-1]
flw.tx_rx.port.tx_name = prt.name

# configure 100 packets to be sent, each having a size of 128 bytes
flw.size.fixed = 128
flw.duration.fixed_packets.packets = 100

# add Ethernet, IP and TCP protocol headers with defaults
flw.packet.ethernet().ipv4().tcp()

# push configuration
api.set_config(config)

# start transmitting configured flows
ts = api.transmit_state()
ts.state = ts.START
api.set_transmit_state(ts)

# fetch & print port metrics
req = api.metrics_request()
req.port.port_names = [prt.name]
print(api.get_metrics(req))
```

## Run test

```Shell
python hello.py
```

## JSON

Every object in snappi can be serialized to or deserialized from a JSON string which conforms to [Open Traffic Generator API](https://github.com/open-traffic-generator/models). This facilitates storing traffic configurations as JSON files and reusing them in API calls with or without further modifications.

* Create a sample config

```python
import snappi

# create a new API instance where location points to controller.
# this will use HTTP transport by default; in order to use gRPC instead,
# one can pass additional kwarg `transport=snappi.Transport.GRPC`:
# api = snappi.api(location="localhost:40051", transport=snappi.Transport.GRPC)
api = snappi.api()
config = api.config()

config.ports.port(name='p1', location='localhost:5555')
config.flows.flow(name='f1')
```

* Serialize to JSON (or python dictionary or YAML)

```python
json_str = config.serialize()
# serialize child of config object to JSON string
json_str = config.ports.serialize()

yaml_str = config.serialize(encoding=config.YAML)
obj_dict = config.serialize(encoding=config.DICT)
```

* Deserialize from JSON (or python dictionary or YAML)

```python
# whether the argument is JSON or YAML or dict is automatically determined
config.deserialize('{"ports": [{"name": "p2", "location": "localhost:5556"}]}')
# deserialize child of config object from JSON string
config.flows.deserialize('[{"name": "f1"}]')

config.deserialize({"ports": [{"name": "p1", "location": "localhost:5555"}]})
config.deserialize('ports:\n- name: p1\n  location: localhost:5555\n')
```

* Pass either snappi object or equivalent JSON string as argument to API calls

```python
config = api.config()

config.ports.port(name='p1', location='localhost:5555')
# config will be serialized to JSON string and sent on wire
api.set_config(config)

json_str = '{"ports": [{"name": "p1", "location": "localhost:5555"}]}'
# JSON string will be directly sent on wire
api.set_config(json_str)
```

Following sections discuss most commonly used constructs in snappi comparing each one of them with equivalent JSON snippet.  
For brevity, snippet for config creation is not included (since it's the same across all).

## Flows

This section deals with flow configuration and control.

### Unidirectional Flow

<details>
<summary>
A simple unidirectional flow for a **one-arm** test.
</summary>

<table>
<tr>
<th>
snappi
</th>
<th>
json
</th>
</tr>
<tr>
<td>

```python
p1 = config.ports.port(name='p1', \
  location='localhost:5555')[-1]
f1 = config.flows.flow(name='f1')[-1]

f1.tx_rx.port.tx_name = p1.name
```

</td>
<td>

```json
{
  "ports": [
    {
      "location": "localhost:5555",
      "name": "p1"
    }
  ],
  "flows": [
    {
      "name": "f1",
      "tx_rx": {
        "port": {
          "tx_name": "p1"
        },
        "choice": "port"
      }
    }
  ]
}
```

</td>
</tr>
</table>
</details>

### Bidirectional Flows

<details>
<summary>A bi-directional flow between two ports.</summary>

<table>
<tr><th>snappi</th><th>json</th></tr><tr>
<td>

```python
p1, p2 = ( \
    config.ports \
    .port(name='p1', location='localhost:5555') \
    .port(name='p2', location='localhost:5556')
)
f1, f2 = config.flows.flow(name='flow p1->p2'). \
  flow(name='flow p2->p1')

f1.tx_rx.port.tx_name = p1.name
f1.tx_rx.port.rx_name = p2.name
f2.tx_rx.port.tx_name = p2.name
f2.tx_rx.port.rx_name = p1.name
```

</td>
<td>

```json
{
  "ports": [
    {
      "location": "localhost:5555",
      "name": "p1"
    },
    {
      "location": "localhost:5556",
      "name": "p2"
    }
  ],
  "flows": [
    {
      "name": "flow p1->p2",
      "tx_rx": {
        "port": {
          "tx_name": "p1",
          "rx_name": "p2"
        },
        "choice": "port"
      }
    },
    {
      "name": "flow p2->p1",
      "tx_rx": {
        "port": {
          "tx_name": "p2",
          "rx_name": "p1"
        },
        "choice": "port"
      }
    }
  ]
}
```

</td>
</tr></table>
</details>

### Meshed Flows

<details>
<summary>Fully meshed flows between four ports.  Each port sends flows to all the ports (except itself).  This example is for four ports, it can be easily extended to an arbitrary number of ports.</summary>

<table>
<tr><th>snappi</th><th>json</th></tr><tr>
<td>

```python
import itertools

for i in range(1, 4):
  config.ports.port(name='p%d' % i, \
    location='localhost:%d' % (5554 + i))

for tx, rx in \
  itertools.permutations([p.name for \ 
  p in config.ports], 2):
  f = config.flows.flow(name='flow %s->%s' \ 
    % (tx, rx))[-1]
  f.tx_rx.port.tx_name = tx
  f.tx_rx.port.rx_name = rx
```

</td>
<td>

```json
{
  "ports": [
    {
      "location": "localhost:5555",
      "name": "p1"
    },
    {
      "location": "localhost:5556",
      "name": "p2"
    },
    {
      "location": "localhost:5557",
      "name": "p3"
    }
  ],
  "flows": [
    {
      "name": "flow p1->p2",
      "tx_rx": {
        "port": {
          "tx_name": "p1",
          "rx_name": "p2"
        },
        "choice": "port"
      }
    },
    {
      "name": "flow p1->p3",
      "tx_rx": {
        "port": {
          "tx_name": "p1",
          "rx_name": "p3"
        },
        "choice": "port"
      }
    },
    {
      "name": "flow p2->p1",
      "tx_rx": {
        "port": {
          "tx_name": "p2",
          "rx_name": "p1"
        },
        "choice": "port"
      }
    },
    {
      "name": "flow p2->p3",
      "tx_rx": {
        "port": {
          "tx_name": "p2",
          "rx_name": "p3"
        },
        "choice": "port"
      }
    },
    {
      "name": "flow p3->p1",
      "tx_rx": {
        "port": {
          "tx_name": "p3",
          "rx_name": "p1"
        },
        "choice": "port"
      }
    },
    {
      "name": "flow p3->p2",
      "tx_rx": {
        "port": {
          "tx_name": "p3",
          "rx_name": "p2"
        },
        "choice": "port"
      }
    }
  ]
}
```

</td>
</tr></table>
</details>

### Protocol Headers With Fixed Fields

<details>
<summary>Simple flow with Ethernet, IP and TCP protocol headers.</summary>
<table>
<tr><th>snappi</th><th>json</th></tr><tr>
<td>

```python
p1 = config.ports.port(name='p1', \ 
  location='localhost:5555')[-1]
f1 = config.flows.flow(name='f1')[-1]

f1.tx_rx.port.tx_name = p1.name
eth, ip, tcp = f1.packet.ethernet().ipv4().tcp()

eth.dst.value = '00:00:00:00:00:AA'
ip.dst.value = '192.168.1.1'
tcp.dst_port.value = 5000
```

</td>
<td>

```json
{
  "ports": [
    {
      "location": "localhost:5555",
      "name": "p1"
    }
  ],
  "flows": [
    {
      "name": "f1",
      "tx_rx": {
        "port": {
          "tx_name": "p1"
        },
        "choice": "port"
      },
      "packet": [
        {
          "ethernet": {
            "dst": {
              "value": "00:00:00:00:00:AA",
              "choice": "value"
            }
          },
          "choice": "ethernet"
        },
        {
          "ipv4": {
            "dst": {
              "value": "192.168.1.1",
              "choice": "value"
            }
          },
          "choice": "ipv4"
        },
        {
          "tcp": {
            "dst_port": {
              "value": 5000,
              "choice": "value"
            }
          },
          "choice": "tcp"
        }
      ]
    }
  ]
}
```

</td>
</tr></table>
</details>

### Protocol Headers With Varying Fields

<details>
<summary>Flow with Ethernet, IP and TCP headers.  Ethernet destination MAC address, destination IP address and TCP destination port are varied using patterns.</summary>
<table>
<tr><th>snappi</th><th>json</th></tr><tr>
<td>

```python
p1 = config.ports.port(name='p1', \ 
  location='localhost:5555')[-1]
f1 = config.flows.flow(name='f1')[-1]

f1.tx_rx.port.tx_name = p1.name
eth, ip, tcp = f1.packet.ethernet().ipv4().tcp()

eth.src.value = '00:00:00:00:00:AA'
eth.dst.values = ['00:00:00:00:00:AB', \ 
  '00:00:00:00:00:AC']

ip.src.value = '192.168.1.1'
ip.dst.increment.start = '192.168.1.2'
ip.dst.increment.step = '0.0.0.1'
ip.dst.increment.count = 2

tcp.src_port.value = 5000
tcp.dst_port.decrement.start = 5002
tcp.dst_port.decrement.step = 1
tcp.dst_port.decrement.count = 2
tcp.seq_num.values = [1, 2]
```

</td>
<td>

```json
{
  "ports": [
    {
      "location": "localhost:5555",
      "name": "p1"
    }
  ],
  "flows": [
    {
      "name": "f1",
      "tx_rx": {
        "port": {
          "tx_name": "p1"
        },
        "choice": "port"
      },
      "packet": [
        {
          "ethernet": {
            "src": {
              "value": "00:00:00:00:00:AA",
              "choice": "value"
            },
            "dst": {
              "values": [
                "00:00:00:00:00:AB",
                "00:00:00:00:00:AC"
              ],
              "choice": "values"
            }
          },
          "choice": "ethernet"
        },
        {
          "ipv4": {
            "src": {
              "value": "192.168.1.1",
              "choice": "value"
            },
            "dst": {
              "increment": {
                "start": "192.168.1.2",
                "step": "0.0.0.1",
                "count": 2
              },
              "choice": "increment"
            }
          },
          "choice": "ipv4"
        },
        {
          "tcp": {
            "src_port": {
              "value": 5000,
              "choice": "value"
            },
            "dst_port": {
              "decrement": {
                "start": 5002,
                "step": 1,
                "count": 2
              },
              "choice": "decrement"
            },
            "seq_num": {
              "values": [
                1,
                2
              ],
              "choice": "values"
            }
          },
          "choice": "tcp"
        }
      ]
    }
  ]
}
```

</td>
</tr></table>
</details>

### Start Flow Transmit

<details>
<summary>Start transmit on a certain set of flows.</summary>
<table>
<tr><th>snappi</th><th>json</th></tr><tr>
<td>

```python
ts = api.transmit_state()
ts.state = ts.START
ts.flow_names = ['f1', 'f2']

api.set_transmit_state(ts)
```

</td>
<td>

```json
{
  "flow_names": [
    "f1",
    "f2"
  ],
  "state": "start"
}
```

</td>
</tr></table>
</details>

## Capture

Capture configuration and control

### Capture Configuration

<details>
<summary>Configure capture prior to starting capture.</summary>
TBD
</details>

### Start Capture

<details>
<summary>Start capture on a set of ports.</summary>
<table>
<tr><th>snappi</th><th>json</th></tr><tr>
<td>

```python
cs = api.capture_state()
cs.state = ts.START
cs.port_names = ['p1', 'p2']

api.set_capture_state(cs)
```

</td>
<td>

```json
{
  "port_names": [
    "p1",
    "p2"
  ],
  "state": "start"
}
```

</td>
</tr></table>
</details>

### Get Capture

<details>
<summary>Retrieve capture for a given port.  Save capture to a .pcap file (python only).</summary><table>
<tr><th>snappi</th><th>json</th></tr><tr>
<td>

```python
req = api.capture_request()
req.port_name = 'p1'

with open('capture.pcap', 'w') as pcap:
  pcap.write(api.get_capture(req).read())
```

</td>
<td>

```json
{
  "port_name": "p1"
}
```

</td>
</tr></table>
</details>

## Metrics

### Port Metrics

<details>
<summary>Get port statistics for a given set of ports.</summary>
<table>
<tr><th>snappi</th><th>json</th></tr><tr>
<td>

```python
req = api.metrics_request()
req.port.port_names = ['tx', 'rx']
req.port.column_names = [req.port.FRAMES_TX, \
  req.port.FRAMES_RX]

res = api.get_metrics(req)
assert res[0].frames_tx == res[1].frames_rx
```

</td>
<td>

```json
{
  "port": {
    "port_names": [
      "p1",
      "p2"
    ],
    "column_names": [
      "frames_tx",
      "frames_rx"
    ]
  },
  "choice": "port"
}
```

</td>
</tr></table>
</details>

### Flow Metrics

<details>
<summary>Get flow statistics.</summary>
TBD
</details>
