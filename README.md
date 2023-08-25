# meta-marketing-api-extraction
Script to extract historical data from the meta (facebook) marketing api and save it to object storage

## Prerequisites  

 - Facebook App, Ad Account and Access Token with the following permissions
    - pages_show_list
    - ads_management
    - leads_retrieval
    - pages_read_engagement
    - pages_manage_metadata
    - pages_manage_ads
 - (gcloud CLI)[https://cloud.google.com/sdk/docs/install] 
 - (Go)[https://go.dev/doc/install]

## What it does  



## Configuration  

Create and `.env` file in the root directory with the following config:
```
ACCOUNT_ID=<AD Account ID>
ACCESS_TOKEN=<Access Token>
BUCKET_NAME=<GCS Bucket to save the data to>
```

## How to run
After cloning the repo, move into the src folder, install the required Go packages and run the script.  
```
cd src
go mod tidy
go run .
```

## Schema
