# Derivatives command description

## Start network with pricefeed oracle



You need to set up oracle for feeding price as a trading price in the `derivatives` module.

First, you set up to start network with the proper config of `pricefeed` in genesis.json using this script, `./start.sh`.

```powershell
./start.sh
```

And after that, in the ununifi, we currently use this for it,
https://github.com/UnUniFi/utils/tree/main/projects/pricefeed
to run pricefeed oracle.   
To run above program, the setting is required as follows,

```powershell
mkdir pricefeed
cd pricefeed
curl -O https://raw.githubusercontent.com/ununifi/utils/main/projects/pricefeed/docker-compose.yml
curl -L https://raw.githubusercontent.com/ununifi/utils/main/projects/pricefeed/launch/[chain-id]/.env.example > .env
cp ./.env
docker-compose up -d
```

Please check `./.env` file for the setting closely.
