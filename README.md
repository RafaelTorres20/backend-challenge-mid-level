# Teste de vaga backend Ideal CVTM

## Rodando o projeto

#### Para rodar o projeto tenha certeza de que as dependências foram devidamente instaladas:

##### Docker

##### Golang 1.2X

1. Adicione as variáveis de ambiente num arquivo .env como mostrado no exemplo `.env.example` não esqueca de utilizar sua chave de api do yahoo finance, você a encontra aqui `https://financeapi.net/dashboard` depois de ter logado na sua conta

2. Criando o banco de dados relacional postgres use o comando no seu terminal:
   ```bash
   docker-compose up -d
   ```
3. Para executar o servidor use o comando
   ```bash
   make run
   ```

### Pronto! agora seu projeto já deve estar rodando

Agora para testar o projeto eu recomendo utilizar o postman, pois o projeto possúi uma collection de requisições prontas para importar no postman!

### instruções de uso da api e como obter cada uma das informações que é esperado no teste!

1. Você deve usar a rota `POST - /users/` para criar o seu usuário, ele é um email e senha, sem nenhum tipo de validaões, nós usaremos ele para fins de conseguirmos utilizar o conceito de autorização do JWT

2. Após criar o usuário, você deve chamar a rota `POST - /users/login/` para que você consiga gerar um token da nossa API para utilizar nos headers de todas as outras requisições exceto as duas supramente mencionadas, você deve conseguir um token, e um id para usar nos testes das outras requisições.

3. Para que o usuário consiga adicionar um ativo a sua lista de ativos, primeiro devemos adicionar o ativo no banco de dados usando a rota `POST - /assets/` Serve para passar um symbol como `GOGL`, `AAPL` ou `MGLU3.SA` por exemplo. Depois disso você pode chamar a rota `POST - /assets/users/{id}`para criar o vínculo de um ativo ao seu usuário.

4. Para ordenar os ativos você deve chamar a rota `POST - /assets/users/{id}/order` passando um json no body da requisição, você pode usar as strings `alpha`, `custom`, `lessPrice` e `greaterPrice` para ordenar respectivamente: ordem alfabética, ordem customizada pelo usuário, ordem de menor valor para o maior e porfim, ordem de maior valor para o menor.

5. Para consultar sua lista de ativos use a rota `GET - /assets/users/{id}` passando seu id você lista todos os ativos que você pode ou não ter ordenado de forma customizada no step anterior.

6. Para consultar os valores atuais dos ativos que você desejar, basta chamar a rota `POST - /assets/prices` ela aceita uma lista de um ou mais sybolos e retorna os ativos com seus preços e suas respectivas moedas.
