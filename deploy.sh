#!/bin/bash

# Source the configuration
if [ -f "./deploy_config.sh" ]; then
    source ./deploy_config.sh
else
    echo "Error: deploy_config.sh not found. Please create this file with your deployment configurations."
    exit 1
fi

# Validate required variables
required_vars=("EC2_HOST" "EC2_APP_DIR" "LOCAL_APP_DIR" "SSH_KEY" "APP_NAME" "AWS_REGION")
for var in "${required_vars[@]}"; do
    if [ -z "${!var}" ]; then
        echo "Error: $var is not set. Please check your deploy_config.sh file."
        exit 1
    fi
done

# Ensure the script is run from the project root
cd "${LOCAL_APP_DIR}" || exit

# Pull latest changes from git
echo "Pulling latest changes from git..."
git pull origin main

# Sync the application to EC2
echo "Syncing application to EC2..."
rsync -avz -e "ssh -i ${SSH_KEY}" \
    --exclude '.git' \
    --exclude 'node_modules' \
    --exclude 'deploy_config.sh' \
    --delete \
    "${LOCAL_APP_DIR}/" "${EC2_HOST}:${EC2_APP_DIR}"

# SSH into the EC2 instance to rebuild and restart the application
echo "Rebuilding and restarting application on EC2..."
ssh -i "${SSH_KEY}" "${EC2_HOST}" << EOF
    cd "${EC2_APP_DIR}"
    go build -o ${APP_NAME} cmd/main.go
    pkill ${APP_NAME} || true
    export GIN_MODE=release
    export AWS_REGION=${AWS_REGION}
    source .env
    nohup ./${APP_NAME} > app.log 2>&1 &
    echo "Application restarted in release mode"
    sleep 2
    if pgrep -f ${APP_NAME} > /dev/null
    then
        echo "Application is running"
    else
        echo "Error: Application failed to start. Check app.log for details."
        exit 1
    fi
EOF

if [ $? -eq 0 ]; then
    echo "Deployment completed successfully"
else
    echo "Deployment failed. Check the output above for errors."
    exit 1
fi