#EXAMPLE GO GRPC

##IDEIA
<br/>
A ideia deste projeto e ter um exemplo de inclusão de pessoa utilizando o protocolo RPC e a 
lib [gRPC] 
(https://grpc.io/) da google. junto fazer um mutiplexador para receber requisiçoes em HTTP tambem.

##OBJETIVO
<br/>
O objetivo deste e estudar e aprender a linguagem o Go e o gRPC. para projetos fulturos.

##TUTORIAL
Estou escrevendo um tutorial de como foi criado a aplicação, para ficar registrado como base 
de consulta. posteriormente pretendo escrever um artigo explicando tudo

(EM CONSTRUÇÃO .............)

###Start Server
1 -  fazer o build do server  
- acessar a pasta "cmd/server" 
- `cd cmd/server`
- fazer o build
- `go build .`
<br>
<br>
2 -  Start o server
- iniciar o servicor co o seguinte comando
- `server.exe -grpc-port={PORTA DO SERVIDOR GRPC} -db-host={HOST DO BD : PORTA} -db-user={USER BD} -db-password={SENHA BD} -db-schema={SCHEMA BD}`
- substituir os {} pelo valor real dentro tem uma descrição do campo

###Start Cliente

1 -  fazer o build do server  
- acessar a pasta "cmd/client-grpc" 
- `cd cmd/client-grpc`
- fazer o build
- `go build .`
<br>
<br>
2 -  Start o Cliente
- iniciar o servicor co o seguinte comando
- `client-grpc.exe -server=localhost:{PORTA DO SERVIDOR GRPC}`
- substituir os {} pelo valor real dentro tem uma descrição do campo


###Bibliografia