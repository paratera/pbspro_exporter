# pbspro_exporter

**This is not an official Paratera product**

Prometheus exporter of PBSPro

## 1.How to build

## 1.1.docker build

```bash
# cd pbspro_exporter/docker
# docker build -t pbspro_exporter:latest .
```
You will get a pbspro_export:latest docker image.

```bash
# docker images
REPOSITORY                     TAG                 IMAGE ID            CREATED             SIZE
pbspro_exporter                latest              db2491b8eda5        7 minutes ago       216MB
```

## 2.How to use pbspro_exporter

### 2.1.docker

```bash
# docker run --name pbspro_exporter -e PBS_ADDR=192.168.100.10 -e EXPORTER_PORT=9107 -d taylor840326/pbspro_exporter:latest
# curl localhost:9107/metrics
```

### Donate

-----

If you like the project and want to buy me a cola, you can through:

| PayPal                                                                                                               | 微信                                                                 |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------- |
| [![](https://www.paypalobjects.com/webstatic/paypalme/images/pp_logo_small.png)](https://www.paypal.me/taylor840326) | ![](https://github.com/taylor840326/blog/raw/master/imgs/weixin.png) |
