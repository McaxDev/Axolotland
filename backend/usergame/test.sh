docker build -t mcaxdev/usergame . &&
docker-compose -f /srv/usergame/docker-compose.yml up -d &&
docker logs -f usergame
