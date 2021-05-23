# Terra Amino Decoder

## How to use

* Start rest server
```
$ make install
$ amino-decoder start
```
* Query with rest server
```
curl -X GET http://127.0.0.1:3000/version
curl -X POST http://127.0.0.1:3000/decode/tx -d '{"amino_encoded_tx": "ENCODED_TX_STRING"}'
curl -X POST http://127.0.0.1:3000/decode/batch -d '{"amino_encoded_tx": ["ENCODED_TX_STRING1", "ENCODED_TX_STRING2]}'
```

* Directly decode amino encoded tx
```
$ amino-decoder version
$ amino-decoder decode tx [amino-encoded-tx]
```

## How to build
```
# it will create build folder and generate binary for each platform (Windows, Mac, Linux)
$ make  
```

## Use docker
```
# docker run --rm -p 6969:3000 -it rmdec/amino-decoder:v1.1.0 amino-decoder start
```
