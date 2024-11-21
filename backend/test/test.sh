docker build -t mcaxdev/test . &&
docker-compose -f /srv/test/docker-compose.yml up -d &&
docker logs -f test
