docker build -t mcaxdev/auth . &&
docker-compose -f /srv/mcaxdev/auth-docker-compose.yml up -d &&
docker logs -f auth
