# Olá

Você pode encontrar abaixo a documentação deste teste, conforme as atividaes solicitadas.
> Além do que foi proposto, eventualmente haverão também alguns comentários sobre esta jornada.

## Para configurar seu repositório

- [x] Realize a substituição de todas as strings `testing/sre-test-1` por `SEU_USUARIO_GIT/NOME_DO_SEU_REPOSITÓRIO` criando um script para fazer essa tarefa (na linguagem de sua escolha) em todos os arquivos.
- [x] Faça o commit e push da alteração para seu repositório.

> Foi criado um script em ShellScript e realizado o commit e push das alterações, já com os arquivos alterados. Mais detalhes sobre esse script, pode ser encontrado no arquivo.

```sh
$ ./replace.sh numeroSette/SRE-TEST-7
Files replaced:
./cmd/get-random-number/register/register.go:   getrandomnumber "github.com/numeroSette/SRE-TEST-7/cmd/get-random-number"
./cmd/get-random-number/register/register.go:   "github.com/numeroSette/SRE-TEST-7/internal/router"
./cmd/main.go:  _ "github.com/numeroSette/SRE-TEST-7/cmd/get-random-number/register"
./cmd/main.go:  "github.com/numeroSette/SRE-TEST-7/internal/config"
./cmd/main.go:  "github.com/numeroSette/SRE-TEST-7/internal/router"
./go.mod:module github.com/numeroSette/SRE-TEST-7
```

## To fix

- [ ] Aplicação não está realizando build da imagem Docker via pipeline no GitHub Actions.
- [ ] Não temos logs no pipeline ou alertas indicando sucesso do teste funcional.
- [ ] Existe um step no pipeline em que realizamos um teste funcional realizando o request para http://localhost:8080/random-number e validamos a resposta, verificar se o teste feito aqui realmente garante que o endpoint está respondendo devidamente.
- [ ] Criar o mesmo teste funcional para a rota `/metrics` da porta **9090**.

## To do

- [ ] Realizar testes de performance na geração de números randômicos.
- [ ] Trazer relatórios sobre estatísticas e métricas dos testes de performance.
- [ ] Diminuir tempo de geração de número randômico.
- [ ] Criar documentação para outros colaboradores contribuírem com o projeto.
- [ ] Implementar métricas sobre o serviço http que responde na rota `/get-random-number` (dicas https://www.robustperception.io/prometheus-middleware-for-gorilla-mux e para uma implementação mais simples, utilize o arquivo [internal/router/router.go](../../internal/router/router.go)) expondo através da rota `/metrics` as métricas adicionais.
- [ ] Reduzir tempo de execução do workflow (GitHub Action).
