# Lola data source
- create a new data source
- call it "lola"
- give it an "InfluxDB" type
- associate it with with http://influxdb:8086 via the Grafana proxy

# Import dashboards
- load dashboard by pasting the corresponding JSON
- associate it with the "lola" data source

------ 

Now, apparently there's a way to let the grafana server to load the lola data source and the associated json dashboards by stashing them into `/etc/grafana/provisioning/datasources` and
`/etc/grafana/provisioning/dashboards` respectively.

Instead of the laborious update by hand we should do something like:
```bash
install -m 0640 -o root -g grafana lola-dashboards.yaml /etc/grafana/provisioning/dashboards
install -m 0640 -o root -g grafana lola-datasource.yaml /etc/grafana/provisioning/datasources

install -m 0755 -o grafana -g grafana /var/lib/grafana/dashboards
install -m 0644 -o grafana -g grafana ...json /var/lib/grafana/dashboards
```
or, maybe, mount the local files at the right locations in the grafana container.
