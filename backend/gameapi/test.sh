docker build -t mcaxdev/gameapi . &&
docker-compose -f /srv/gameapi/docker-compose.yml up -d &&
docker logs -f gameapi
