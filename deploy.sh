gcloud functions deploy metapass-images-function-rinkeby --entry-point TokenMetadata --runtime go113 --trigger-http --allow-unauthenticated --update-env-vars NODE_URL=https://rinkeby.infura.io/v3/31a18619fa0c4165a463cbd0fd7ba064,CONTRACT_ADDRESS=0x5661460F87ecceA19BB3DF10703F855FF7DbF860