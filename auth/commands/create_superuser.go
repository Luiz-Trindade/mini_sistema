package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"mini_sistema/auth/services"

	"golang.org/x/term"
)

type SuperuserCommand struct {
	userService *services.UserService
}

func NewSuperuserCommand(userService *services.UserService) *SuperuserCommand {
	return &SuperuserCommand{userService: userService}
}

func (c *SuperuserCommand) Run() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Criar Superusuário ===")

	fmt.Print("Nome: ")
	nomeInput, _ := reader.ReadString('\n')
	nome := strings.TrimSpace(nomeInput)

	fmt.Print("E-mail: ")
	emailInput, _ := reader.ReadString('\n')
	email := strings.TrimSpace(emailInput)

	senha, err := c.readPassword("Senha: ")
	if err != nil {
		log.Fatalf("\nErro ao ler a senha: %v\n", err)
	}

	senhaConfirmacao, err := c.readPassword("Confirme a Senha: ")
	if err != nil {
		log.Fatalf("\nErro ao ler a confirmação da senha: %v\n", err)
	}

	if senha == "" {
		log.Fatal("\nErro: A senha não pode ser vazia.")
	}

	if senha != senhaConfirmacao {
		log.Fatal("\nErro: As senhas não coincidem.")
	}

	user, err := c.userService.CreateSuperuser(nome, email, senha)
	if err != nil {
		log.Fatalf("\nErro ao criar superusuário: %v\n", err)
	}

	fmt.Printf("\nSuperusuário '%s' (%s) criado com sucesso!\n", user.Nome, user.Email)
}

func (c *SuperuserCommand) readPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	fmt.Println()
	return strings.TrimSpace(string(bytePassword)), nil
}
