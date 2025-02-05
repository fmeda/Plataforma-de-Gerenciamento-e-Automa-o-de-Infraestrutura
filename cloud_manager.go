package main

import (
    "fmt"
    "os/exec"
    "log"
    "os"
)

func runCommand(command string, args []string) {
    cmd := exec.Command(command, args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        log.Fatalf("Erro ao executar o comando: %v\n", err)
    }
}

func createInfrastructure() {
    fmt.Println("Criando infraestrutura com Terraform...")
    runCommand("terraform", []string{"init"})
    runCommand("terraform", []string{"apply", "-auto-approve"})
}

func main() {
    fmt.Println("Gerenciamento de Infraestrutura iniciado.")
    createInfrastructure()
}
