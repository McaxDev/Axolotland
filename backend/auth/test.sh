docker build -t mcaxdev/auth . &&
docker-compose -f /srv/auth/docker-compose.yml up -d &&
docker logs -f auth
