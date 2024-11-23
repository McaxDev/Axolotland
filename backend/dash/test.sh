docker build -t mcaxdev/dash . &&
docker-compose -f /srv/mcaxdev/dash-docker-compose.yml up -d &&
docker logs -f dash
