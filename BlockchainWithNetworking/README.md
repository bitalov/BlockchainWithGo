## Basic TCP Broadcasting Blockchain !

Different blockchains can communicate with each other

### TCP Broadcasting Illustration

![illustration+Blockchain](https://user-images.githubusercontent.com/60072763/164713618-2c15c65a-4e4d-44e5-9c22-98878c1abd71.png)


### Steps to Deploy What I've implemented so far
- Download *main.go* , and *tcp_ports.env*.
- Rename ports.env to .env using "mv tcp_ports.env .env".
- run main.go using `run go main.go`
- To Add a new Block you can connect to the server "Terminal 1" to send the block you created using `nc localhost 9000` *9000 is the port in the .env file*
- You can open as many terminals as you want and repeat the previous step


### Screenshots


