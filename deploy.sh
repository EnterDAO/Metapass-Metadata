gcloud functions deploy metapass-images-rinkeby --entry-point TokenMetadata --runtime go113 --trigger-http --allow-unauthenticated --update-env-vars NODE_URL=https://rinkeby.infura.io/v3/31a18619fa0c4165a463cbd0fd7ba064,CONTRACT_ADDRESS=0x46c6501E5EDbc5D81ABB45E10B2dfF4Bb8b70553,GENERATOR_ENDPOINT=http://34.152.32.5:8081/generate-video