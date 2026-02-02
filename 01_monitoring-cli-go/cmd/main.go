package main

import (
	"bufio"
	"fmt"
	"io"

	"net/http"
	"os"
	"strings"
	"time"
)

const (
	monitoramentos = 1
	delay          = 0
)

func main() {
	exibiIntroducao()

	for {
		exibiMenu()

		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 3:
			fmt.Println("saindo...")
			os.Exit(0)
		default:
			fmt.Println("não conheço esse comando!!")
			os.Exit(-1)
		}
	}
}

func exibiIntroducao() {
	nome := "monitoring sites cli"
	versao := 1.1

	fmt.Println(nome)
	fmt.Println("version:", versao)
}

func exibiMenu() {
	fmt.Println("1-> iniciar monitoramento")
	fmt.Println("2-> exibir logs")
	fmt.Println("3-> sair do programa")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)

	return comando
}

func iniciarMonitoramento() {
	fmt.Println("monitorando...")
	defer fmt.Println()

	// sites := []string{
	// 	"https://www.alura.com.br",
	// 	"https://httpbin.org/status/404",
	// 	"https://www.google.com",
	// 	"https://httpbin.org/status/200",
	// }

	sites := leSitesDoArquivo()

	for range monitoramentos {
		for i, site := range sites {
			fmt.Println("testando site", i, ":", site)
			testaSite(site)
		}

		time.Sleep(delay * time.Minute)
		fmt.Println()
	}

}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("error:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("site:", site, "está com problemas. Status code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("./docs/sites.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer arquivo.Close()

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("error:", err)
		}
	}

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("./docs/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer arquivo.Close()

	fmt.Fprintf(arquivo, "%s - %s - online: %t\n", time.Now().Format("02/01/2006 15:04:05"), site, status)
}

func imprimeLogs() {
	fmt.Println("exibindo logs...")

	arquivo, err := os.ReadFile("./docs/log.txt")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(arquivo))
	fmt.Println()
}
