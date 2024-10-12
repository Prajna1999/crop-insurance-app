#!/bin/bash

# Configuration
EC2_HOST="user@ec2-xx-xx-xx-xx.compute-1.amazonaws.com"
EC2_APP_DIR="/path/to/app/on/ec2"
LOCAL_APP_DIR="/path/to/local/app"
SSH_KEY="/path/to/your/ec2-key.pem"
APP_NAME="crop-insurance-app"

# Ensure the script is run from the project root
cd "${LOCAL_APP_DIR}"

# Pull latest changes from git
git pull origin main

# Sync the application to EC2
rsync -avz -e "ssh -i ${SSH_KEY}" \
    --exclude '.git' \
    --exclude 'node_modules' \
    --delete \
    "${LOCAL_APP_DIR}/" "${EC2_HOST}:${EC2_APP_DIR}"

# SSH into the EC2 instance to rebuild and restart the application
ssh -i "${SSH_KEY}" "${EC2_HOST}" << EOF
    cd "${EC2_APP_DIR}"
    go build -o ${APP_NAME} cmd/main.go
    pkill ${APP_NAME}
    export GIN_MODE=release
    source .env
    nohup ./${APP_NAME} > app.log 2>&1 &
    echo "Application restarted in release mode"
EOF

echo "Deployment completed"