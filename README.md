# This a simple and basic blockchain using GO Lang !

### TODO
- **Networking** : So far this is a blockchain running only in a signle node , I'll try to add networking parts where different nodes can contribute new blocks !
- **Proof of Work** : I'll try to implement PoW Algorithm.
- **Proof of Stake** : I'll try to implement PoS Algorithm.
- **Peer-to-Peer** : I'll try to make the blockchain P2P.

### Steps to Deploy What I've implemented so far
- Download *main.go* , and *ports.env*.
- Rename ports.env to .env using "mv ports.env .env".
- Open a web browser and `http://localhost:8080/` (8080 is the port I wrote in .env file).
- To Add a new Block you can use curl to send a `POST` request to `http://localhost:8080` .

  **Example** : `curl -i -X POST -H 'Content-Type: application/json' -d '{"BPM":50}' http://localhost:8080`

### Screenshots

*Screen of POST request using curl*

*Screen of Blocks*
