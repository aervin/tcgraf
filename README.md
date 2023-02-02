# tcgraf

##### A service for managing Grafana dashboards

#### Usage

The service accepts a single data structure on port 80:
```json
{
    "type": "init_dash",
    "instance_id": "unique-instance-id",
    "nas_id": "unique-nas-id"
}
```

A `.env` file should be added to the root of the repo with two key/value entries:
```
TC_GRAFANA_ADDRESS=<Grafana address>
TC_GRAFANA_KEY=<Grafana API key>
```

Basic usage example:
```bash
make run
curl -d '{"type": "init_dash", "instance_id": "tc-instance-id-here", "nas_id": "JDrrcq7kw0lgvmpu"}' localhost:9999
```

#### Assumptions
* TrueCommand is exposing NAS stats for a Prometheus instance to scrape.
* A Grafana instance is running at the address provided in the `.env` file. Additionally, a collection of Grafana Library Panels has been prepopulated.

See the `makefile` for development commands.