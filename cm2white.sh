#!/bin/bash
if [[ -d "~/metaplex" ]]
then
    rm -rf ~/metaplex
fi

git clone https://github.com/metaplex-foundation/metaplex.git ~/metaplex
cd ~/metaplex/js
yarn install
yarn bootstrap
yarn build

echo "change devnet to mainnet-beta if you want to publish to the mainnet"

solana-keygen new -o ~/.config/solana/id.json > id
solana-keygen new --outfile ~/.config/solana/devnet.json > devnet
solana config set --keypair ~/.config/solana/devnet.json
solana config set --url https://api.devnet.solana.com/
solana airdrop 2 --url devnet
sleep 5
solana airdrop 2 --url devnet
cat ~/.config/solana/devnet.json
solana-keygen new --outfile ~/.config/solana/whitelist.json > whitelist
cat ~/.config/solana/whitelist.json
spl-token create-token --decimals 0 ~/.config/solana/whitelist.json
echo "enter the pubkey of whitelist.json"
read -sp pubwhite
spl-token create-account $pubwhite
spl-token accounts
echo "enter the number of token for whitelist"
read numwhite
spl-token mint $pubwhite $numwhite
spl-token accounts
echo "enter target whitelist account pubkey"
read tarwhite
spl-token transfer $pubwhite 1 $tarwhite --fund-recipient

ts-node ~/metaplex/js/packages/cli/src/candy-machine-v2-cli.ts upload -e devnet -k ~/.config/solana/devnet.json -cp ./config.json -c temp ./assets
ts-node ~/metaplex/js/packages/cli/src/candy-machine-v2-cli.ts verify_upload -e devnet -k ~/.config/solana/devnet.json -c temp
