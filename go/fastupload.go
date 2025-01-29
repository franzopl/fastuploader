package main

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Por favor, forneça um arquivo ou diretório.")
        os.Exit(1)
    }

    diretorioDestino := fmt.Sprintf("/home/%s/nzb/", os.Getenv("USER"))

    // Ensure the destination directory exists
    if err := os.MkdirAll(diretorioDestino, os.ModePerm); err != nil {
        fmt.Printf("Erro ao criar o diretório de destino: %v\n", err)
        os.Exit(1)
    }

    usePack := false
    obfuscate := false

    // Parse command line options
    for _, arg := range os.Args[1:] {
        switch arg {
        case "-p":
            usePack = true
        case "-x":
            obfuscate = true
        default:
            if !strings.HasPrefix(arg, "-") {
                processInput(arg, usePack, obfuscate, diretorioDestino)
                return
            }
            fmt.Printf("Opção inválida -%s\n", arg[1:])
            os.Exit(1)
        }
    }
}

func processInput(path string, usePack, obfuscate bool, diretorioDestino string) {
    info, err := os.Stat(path)
    if err != nil {
        fmt.Printf("O argumento fornecido não é um arquivo ou diretório válido: %v\n", err)
        os.Exit(1)
    }

    if info.IsDir() {
        processDirectory(path, usePack, obfuscate, diretorioDestino)
    } else {
        processFile(path, usePack, obfuscate, diretorioDestino)
    }
}

func processFile(filePath string, usePack, obfuscate bool, diretorioDestino string) {
    name := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
    nfoFile := filepath.Join(filepath.Dir(filePath), name+".nfo")

    // Generate .nfo file using mediainfo
    if _, err := exec.Command("mediainfo", filePath).Output(); err != nil {
        fmt.Printf("Erro ao gerar o arquivo .nfo: %v\n", err)
        return
    }

    // Move the .nfo file to the specified directory
    if err := os.Rename(nfoFile, filepath.Join(diretorioDestino, filepath.Base(nfoFile))); err != nil {
        fmt.Printf("Erro ao mover o arquivo .nfo: %v\n", err)
        return
    }

    packArgs := []string{"-i", filePath}
    if usePack {
        packArgs = append(packArgs, "--pack")
    } else {
        packArgs = append(packArgs, "--gen_par2")
    }
    if obfuscate {
        packArgs = append(packArgs, "-x")
    }

    // Run ngPost
    if _, err := exec.Command("ngPost", packArgs...).Output(); err != nil {
        fmt.Printf("Erro ao executar ngPost: %v\n", err)
    }
}

func processDirectory(dirPath string, usePack, obfuscate bool, diretorioDestino string) {
    files, err := os.ReadDir(dirPath)
    if err != nil {
        fmt.Printf("Erro ao ler o diretório: %v\n", err)
        return
    }

    for _, file := range files {
        if !file.IsDir() {
            processFile(filepath.Join(dirPath, file.Name()), usePack, obfuscate, diretorioDestino)
        }
    }
}
