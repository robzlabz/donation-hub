#!/bin/bash

echo "Local Deployment..."

# Create s3 bucket for media donation
awslocal s3api create-bucket --bucket media-donation-hub --create-bucket-configuration LocationConstraint=ap-southeast-3
