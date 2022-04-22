# Basic TCP Broadcasting Blockchain

Different blockchains can communicate now with each other

### TCP Broadcasting Illustration :

![illustration](https://user-images.githubusercontent.com/60072763/164738471-e2d42b9c-ed18-4d40-afdb-0252b2a9d39e.png)

### Deployement :
- Download *main.go* , and *tcp_ports.env*.
- Rename ports.env to .env using `mv tcp_ports.env .env`.
- run main.go using `run go main.go`
- To Add a new Block you can connect to the server ***"Terminal 1"*** to send the block you created using `nc localhost 9000` *9000 is the port in the .env file*
- You can open as many terminals as you want and repeat the previous step


### Screenshots :

![deployed_Screen](https://user-images.githubusercontent.com/60072763/164743177-c4e0e47c-a2a4-4fd4-a430-7cec099fcd1b.PNG)
