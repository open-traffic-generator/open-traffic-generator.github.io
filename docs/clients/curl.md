# curl

`curl` command-line utility is capable of communicating with any REST-compatible API endpoint over HTTP(s), and OTG is no different. Although you would not want `curl` to be your go-to choice for OTG, mastering some key queries will be quite useful down the road.

For the examples below to work, start with initializing an environmental variable `OTG_HOST` with a correct URL string for your OTG API Endpoint deployment. Here we assume you've deployed the Endpoint on the same host where you'll run `curl`, over a default HTTPs port – similar to a basic [Ixia-c Traffic Engine deployment](../implementations.md#ixia-c).

```Shell
OTG_HOST="https://localhost"
```

## Apply a configuration

Suppose you have an OTG configuration stored in `otg.json` file. To apply it to an OTG Endpoint, use:

```Shell
curl -k "${OTG_HOST}/config" \
    -H "Content-Type: application/json" \
    -d @otg.json
```

## Show current configuration

Now that you have your configuration applied, you can check how the OTG Endpoint took it. You can find yourself pulling the configuration this way even when developing complex test programs or CI/CD pipelines – to see what ended up being applied to a Traffic Generator.

```Shell
curl -k "${OTG_HOST}/config"
```

For example, if you want to make a small change within a complex configuration that was applied by a Test Program, show current configuration and save it to a file, make changes and apply an updated configuration from the file.

## Start transmitting flows

Now that you have a configuration applied, you can start transmitting all configured Traffic Flows:

```Shell
curl -k "${OTG_HOST}/control/transmit" \
    -H  "Content-Type: application/json" \
    -d '{"state": "start"}'
```

## Stop transmitting flows

Even more important command is to stop transmitting of all configured Traffic Flows. This is especially useful if you started a long-running test and your test program crashed after that.

```Shell
curl -k "${OTG_HOST}/control/transmit" \
    -H  "Content-Type: application/json" \
    -d '{"state": "stop"}'
```
