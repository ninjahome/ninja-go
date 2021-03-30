# ninja-go
peer to peer chat protocol based on blockchain technology.

##compile from source code

    git clone https://github.com/ninjahome/ninja-go.git
    cd ninja-go
    make mac|linux|arm|win

##init node
init node configuration and wallet

    ninja.[mac|lnx|exe] init

create new wallet

    ninja.mac wallet create -p [PWD] [-t]

-  -t: test network
- PWD: password of the wallet
        
output when wallet created:
  
        /.ninja/test_keystore/UTC--2021-03-30T01-14-14.326901000Z--afb2a45a496b53f455225eebe501a23e33a6ddd8058e3ad1fdc5d9b6ee2ee4e0c188bfbaeb72f61865840569601fe1f4
    create success!

##start node
    ninja.mac -n test