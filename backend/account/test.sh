docker build -t mcaxdev/account . &&
docker-compose -f /srv/account/docker-compose.yml up -d &&
docker logs -f account
