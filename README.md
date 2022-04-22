# Basic Blockchain with Golang

### TODOs :
- [x] [**Networking**](https://github.com/bitalov/BlockchainWithGo/tree/main/BlockchainWithNetworking).
- [ ] **Proof of Work**.
- [ ] **Proof of Stake**.
- [ ] **Peer-to-Peer**.

### Deployement :
- Download *main.go* , and *ports.env*.
- Rename ports.env to .env using "mv ports.env .env".
- run main.go using `run go main.go`
- Open a web browser and `http://localhost:8080/` (8080 is the port I wrote in .env file).
- To Add a new Block you can use curl to send a `POST` request to `http://localhost:8080` .

  **Example** : `curl -i -X POST -H 'Content-Type: application/json' -d '{"BPM":50}' http://localhost:8080`

### Screenshots :

#### Screenshot of POST request using "CURL"

![Screen1](https://user-images.githubusercontent.com/60072763/163896882-2ff5bb40-bbce-4763-a02d-c9bd5f7b9c1a.PNG)

#### Screenshot of Blocks

![Screen2](https://user-images.githubusercontent.com/60072763/163896916-c2483bc3-5cfc-4fc1-bfaf-faf8bc34a4e8.PNG)
