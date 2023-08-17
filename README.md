# Periodo de desenvolvimento
- Dias de desenvolvimento: 11/08/2022, 15/08/2022, 16/08/2022 

# Requisitos
- Docker
- Docker Compose (aceite version '2')

# Instruções de instalação #
- Acesse a pasta onde irá clonar o projeto

- Escolha uma pasta de preferência e clone o projeto
```
$ git clone https://github.com/ertfly/napp-test.git
```

- Ou com ssh
```
$ git clone git@github.com:ertfly/napp-test.git
```

- Acesse a pasta do projeto
```
$ cd napp-test
```

- Copie o arquivo **docker-compose.sample.yml** renomeando para **docker-compose.yml**
```
$ cp docker-compose.sample.yml docker-compose.yml
```
> **_NOTA:_**  Os arquivos copiados estão aplicados no .gitignore, e não causará efeitos de modificação

- Copie o arquivo **.env.example** renomeando para **.env**
```
$ cp .env.example .env
```
> **_NOTA:_**  Configure o **.env** caso houver necessidade de alterar a porta da API
> **_NOTA:_**  O docker-compose foi configurado para que os containers tenha o seus hosts utilizando o atributo ***container_name**, exemplo o **DB_HOST** do arquivo **.env** ficaria **DB_HOST=teste.db** o nome dado no atributo, então é necessário que o docker-compose, na versão sitada, suba os containers com os nomes definidos, caso não terá que alterar os dados de acesso do banco, caso OK não precisa alterar pode deixar os dados como estão apenas execute a cópia.

- Criei o network dos containers
```
$ docker network create test-net
```
> **_NOTA:_**  Se a rede test-net já existir ignore.

- Altere o arquivo **docker-compose.yml** substitua na parte **8000:** pela porta web disponível na sua máquina caso houver necessidade
```
    ...
    ports:
      - '8000:8000'
    ...
``` 

- Altere o arquivo **docker-compose.yml** substitua na parte **3306:** pela porta web disponível na sua máquina caso houver necessidade
```
    ...
    ports:
      - '3306:3306'
    ...
``` 

- Uma vez alterado o arquivo **docker-compose.yml** vamos utilizar o docker-compose para criar os containers
```
$ docker-compose up -d --build
```
> **_NOTA:_**  O migration ao executar as API's são rodados automaticamentes no run do build

## O que foi criado
- API de acordo com o teste solicitado.