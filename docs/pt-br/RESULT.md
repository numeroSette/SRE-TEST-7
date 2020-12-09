# Olá

Você pode encontrar abaixo a documentação deste teste, conforme as atividaes solicitadas.
> Além do que foi proposto, eventualmente haverão também alguns comentários sobre esta jornada.

## Para configurar seu repositório

- [x] Realize a substituição de todas as strings `testing/sre-test-1` por `SEU_USUARIO_GIT/NOME_DO_SEU_REPOSITÓRIO` criando um script para fazer essa tarefa (na linguagem de sua escolha) em todos os arquivos.
- [x] Faça o commit e push da alteração para seu repositório.

> Foi criado um script em ShellScript ([replace.sh](../../replace.sh)) e realizado o commit e push das alterações, já com os arquivos alterados. Mais detalhes sobre esse script, pode ser encontrado no arquivo.

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

- [x] Aplicação não está realizando build da imagem Docker via pipeline no GitHub Actions.

> Corrigida a url `docker.pkg.gitbuh.com` por `docker.pkg.github.com` no arquivo [main.yml](../../.github/workflows/main.yml):

```yaml
    - name: Push image
    run: |
        IMAGE_ID=docker.pkg.github.com/${{ github.repository }}/$IMAGE_NAME
```

- [x] Não temos logs no pipeline ou alertas indicando sucesso do teste funcional.

> Substituído o parâmetro `-s` por `-v` no teste feito para validar a resposta do teste funcional, dessa maneira é possível ter uma resposta detalhada sobre o retorno da chamada ao endpoint:

> Alterações foram realizadas no arquivo [main.yml](../../.github/workflows/main.yml):

```yaml
      - name: Test URL response (/get-random-number)
        run: curl -v http://localhost:8080/get-random-number
```

```sh
$ curl -v http://localhost:8080/get-random-number
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> GET /get-random-number HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Sun, 06 Dec 2020 05:54:44 GMT
< Content-Length: 116
< 
{"random_number":"17722303992115235272844924504241253148611320514081232194338332455231136341816473746263514294410"}
* Connection #0 to host localhost left intact
* Closing connection 0
```

> Foi também adicionado o parâmetro `-f` que retorna um código 22 para a maioria das tentativas de chamadas mal-sucedidas, é interessante observar que esse método não é totalmente seguro, porém para atender o cenário atual sem agregar complexidade foi utilizada essa estratégia

```text
-f, --fail

(HTTP) Fail silently (no output at all) on server errors. 

This is mostly done to better enable scripts etc to better deal with failed attempts. 

In normal cases when an HTTP server fails to deliver a document, 
it returns an HTML document stating so (which often also describes why and more). 

This flag will prevent curl from outputting that and return error 22.

This method is not fail-safe and there are occasions where non-successful response codes will slip through, especially when authentication is involved (response codes 401 and 407).
```

> Dessa maneira quando for chamada, a pipeline pode entender que esse retorno trata-se de um erro e falhar no step. Também é possível criar uma condição customizada para aceitar apenas alguns tipos de retornos, porém os mesmos devem previstos.

> Alterações foram realizadas no arquivo [main.yml](../../.github/workflows/main.yml):

```yaml
      - name: Test URL response (/get-random-number)
        run: curl -fv http://localhost:8080/get-random-number
```

```sh
$ curl -fv http://localhost:8080/get-random-number
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> GET /get-random-number HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
* The requested URL returned error: 404 Not Found
* Closing connection 0
curl: (22) The requested URL returned error: 404 Not Found
```

- [x] Existe um step no pipeline em que realizamos um teste funcional realizando o request para http://localhost:8080/random-number e validamos a resposta, verificar se o teste feito aqui realmente garante que o endpoint está respondendo devidamente.

> Alterada rota `/random-number` por `/get-random-number` no arquivo no arquivo [register.go](../../cmd/get-random-number/register/register.go):

```golang
func init() {
    router.Router.HandleFunc("/get-random-number", getrandomnumber.GetRandomNumber)
}
```

> Esta rota não correspondia ao que estava configurado na pipeline, conforme pode ser visto no arquivo [main.yml](../../.github/workflows/main.yml):

```yaml
      - name: Test URL response (/get-random-number)
        run: curl -fv http://localhost:8080/get-random-number
```

- [x] Criar o mesmo teste funcional para a rota `/metrics` da porta **9090**.

> Adicionada a configuração para a exposição de portas (host port e container port) `-p 9090:9090` no arquivo [main.yml](../../.github/workflows/main.yml), esta configuração estava faltando para que o endpoint `/metrics` pudesse ser alcançado

```yaml
      - name: Run service
        run: |
          docker run --rm \
            --name service \
            -d \
            -p 8080:8080 \
            -p 9090:9090 \            
            docker.pkg.github.com/${{ steps.docker_config.outputs.image_repository }}/service:latest
```

> Foi adicionado também no [main.yml](../../.github/workflows/main.yml), o mesmo teste de endpoint que utilizamos anteriormente para `get-random-number`

```yaml
      - name: Test URL response (/metrics)
        run: curl -fv http://localhost:9090/metrics
```

## To do

- [x] Realizar testes de performance na geração de números randômicos.
- [x] Trazer relatórios sobre estatísticas e métricas dos testes de performance.
- [x] Diminuir tempo de geração de número randômico.

> Inicialmente foram realizados dois testes de performance, o primeiro foi realizado através do navegador, conforme pode-se verificar através da imagem abaixo:
![Benchmarking Chrome](./images/GetRandomNumber.png)

> O segundo teste foi realizado por meio da criação de um teste de benchmark que pode ser encontrado no arquivo [function_test.go](../../cmd/get-random-number/function_test.go):

```sh
$ docker exec -i service bash -c "cd cmd/get-random-number && go test -run=Bench -bench=."
goos: linux
goarch: amd64
pkg: github.com/numeroSette/SRE-TEST-7/cmd/get-random-number
BenchmarkGetRandomNumberTest-2                 1        1442997725 ns/op
PASS
ok      github.com/numeroSette/SRE-TEST-7/cmd/get-random-number 1.452s
```
> Foi criado também um endpoint chamado `get-random-number-native`, este endpoint foi criado para que possa haver uma comparação entre a geração dos numeros randômicos, dessa maneira em vez de alterar o anterior foi somado um novo.
> Os arquivos referentes a este novo endpoint podem ser encontrados [neste caminho](../../cmd/get-random-number-native):

```golang
package getrandomnumbernative

import (
  "encoding/json"
  "fmt"
  "log"
  "math/rand"
  "net/http"
  "strings"
  "time"
)

// randomInt returns a string with random numbers
func randomInt() string {

  // Reference
  // https://flaviocopes.com/go-random/
  // https://www.random.org/sequences/?min=1&max=52&col=1&format=plain&rnd=new

  // Use Unix date to seed
  rand.Seed(time.Now().UnixNano())

  // Generate 52 numbers between 1 and 60
  var array = make([]int, 52)

  for i := 0; i < 52; i++ {
    array[i] = 1 + rand.Intn(60-1)
  }

  return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), ""), "[]")
}

// RandomNumberResponse is a struct for response
type RandomNumberResponse struct {
  RandomNumber string `json:"random_number"`
}

// GetRandomNumberNative is a function to return a random number
func GetRandomNumberNative(response http.ResponseWriter, request *http.Request) {

  out := &RandomNumberResponse{
    RandomNumber: randomInt(),
  }

  response.Header().Set("Content-Type", "application/json")

  err := json.NewEncoder(response).Encode(out)
  if err != nil {
    log.Printf("failed to encode json to HTTP response: %v", err)
  }

}

```

> Foi utilizada a mesma estrutura que `get-random-number` para a chamada deste novo endpoint (`get-random-number-native`), assim foi possível realizar os testes de forma igual e comparar apenas a maneira em que o número randômico é gerado nas duas abordagens:

```golang
package getrandomnumbernative

import (
  "net/http"
  "net/http/httptest"
  "regexp"
  "testing"
)

func GetRandomNumberNativeTest(b *testing.B) {

  // References
  // https://blog.questionable.services/article/testing-http-handlers-go/

  // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
  // pass 'nil' as the third parameter.
  req, err := http.NewRequest("GET", "/get-random-number-native", nil)
  if err != nil {
    b.Fatal(err)
  }

  // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(GetRandomNumberNative)

  // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
  // directly and pass in our Request and ResponseRecorder.
  handler.ServeHTTP(rr, req)

  // Check the status code is what we expect.
  if status := rr.Code; status != http.StatusOK {
    b.Errorf("handler returned wrong status code: got %v want %v",
      status, http.StatusOK)
  }

  // Check the response body is what we expect.
  expected := `^\{\"random_number"\:\"\d+\"\}$`
  matched, _ := regexp.MatchString(`^\{\"random_number"\:\"\d+\"\}$`, rr.Body.String())
  if matched {
    b.Errorf("handler returned unexpected body: got %v want %v",
      rr.Body.String(), expected)
  }
}

func BenchmarkGetRandomNumberNativeTest(b *testing.B) { GetRandomNumberNativeTest(b) }

```

> O Resultados do teste com o novo endpoint podem ser verificados abaixo:
![Benchmarking Chrome](./images/GetRandomNumberNative.png)

```sh
$ docker exec -i service bash -c "cd cmd/get-random-number-native && go test -run=Bench -bench=."
goos: linux
goarch: amd64
pkg: github.com/numeroSette/SRE-TEST-7/cmd/get-random-number-native
BenchmarkGetRandomNumberNativeTest-2    1000000000               0.000058 ns/op
PASS
ok      github.com/numeroSette/SRE-TEST-7/cmd/get-random-number-native  0.020s
```

> Os testes de benchmarking também estão sendo realizados na pipeline ([main.yml](../../.github/workflows/main.yml))

- [ ] Criar documentação para outros colaboradores contribuírem com o projeto.
- [ ] Implementar métricas sobre o serviço http que responde na rota `/get-random-number` (dicas https://www.robustperception.io/prometheus-middleware-for-gorilla-mux e para uma implementação mais simples, utilize o arquivo [internal/router/router.go](../../internal/router/router.go)) expondo através da rota `/metrics` as métricas adicionais.
- [ ] Reduzir tempo de execução do workflow (GitHub Action).
