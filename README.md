This Project is to go through Prometheus functionalities

docker build -t promtest . && \
docker rm -f promtest; \
docker run -dp 8081:8081 --name promtest promtest && \
docker logs -f promtest


---
docker run -p 9090:9090 -v /prometheus/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus