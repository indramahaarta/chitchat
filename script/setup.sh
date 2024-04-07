#!/bin/bash

echo "Start to execute script"
echo "============================================================"

copy_env() {
    local dir=$1
    local env_file_name=${2:-".env"}

    cp .env "./$dir/$env_file_name"
    echo "Environment file copied to $dir/$env_file_name"
    echo "'$dir': Done"
    echo "============================================================"
}

declare -a repos=( 
    "frontend|.env"
    "backend|app.env"
)

for repo in "${repos[@]}"; do
    IFS='|' read -r dir env_file_name <<< "$repo"
    copy_env "$dir" "$env_file_name"
done

if command -v docker >/dev/null 2>&1; then
    echo "Docker is installed."
else
    echo "Docker is not installed. Please install :)"
    exit 1
fi

if command -v docker-compose >/dev/null 2>&1; then
    echo "Docker Compose is installed."
else
    echo "Docker Compose is not installed. Please install; :)"
    exit 1
fi

docker-compose down
docker-compose up -d --build
docker image prune -f