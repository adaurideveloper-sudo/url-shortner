# Encurtador de URL em Go

Este projeto é um **encurtador de URL** simples feito em Go, usando apenas memória (não salva os dados em banco).

## Como funciona?

Você pode enviar uma URL original, e o serviço irá criar um link curto que redireciona para a URL informada.

- As URLs são guardadas apenas na memória durante o funcionamento do servidor.
- O servidor oferece:
  - Um endpoint para encurtar URLs (`/encurtar`)
  - Um endpoint para redirecionar acessos (`/SEUCODIGO`)

## Como usar

### 1. Clone o projeto

```bash
git clone https://github.com/SEUUSUARIO/SEUPROJETO.git
cd SEUPROJETO
```

### 2. Execute o servidor

```bash
go run main.go
```

O servidor estará rodando em: [http://localhost:8080](http://localhost:8080)

---

### 3. Gerando uma URL curta

- No navegador ou pelo terminal (curl):

```bash
curl "http://localhost:8080/encurtar?url=https://google.com"
```

- Você também pode informar um código customizado:

```bash
curl "http://localhost:8080/encurtar?url=https://google.com&custom=meucodigo"
```

---

### 4. Usando o link curto

Após encurtar, o serviço irá mostrar o link gerado, exemplo:

```
Link criado com sucesso!
Acesse: http://localhost:8080/abcde
```

Acesse esse link para ser redirecionado para a URL original.

---

## Observações importantes

- **Os dados são temporários**: ao parar o servidor, todos os links encurtados desaparecem.
- **Para produção**: seria necessário usar um banco de dados real e adicionar mais validações.

## Como funciona o código?

- O servidor usa um mapa em memória para guardar as URLs.
- Usa um `RWMutex` para garantir que o mapa não seja acessado simultaneamente por várias requisições (evitando erros).
- O endpoint `/encurtar` recebe a URL original e gera um código curto.
- O endpoint `/SEUCODIGO` redireciona para o endereço original se o código existir.

---

## Requisitos

- [Go](https://golang.org/) na versão 1.20 ou superior.

## Contribuições

Pull requests são bem-vindos! Se tiver sugestões, abra um issue.

---

## Licença

MIT
