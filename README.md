# Go_modelo

## 🚀 Rodando em modo desenvolvimento

### 1. Instalar o Go

Certifique-se de ter o Go instalado:  
https://go.dev/dl/

### 2. Instalar o Air (hot reload)

bash
go install github.com/air-verse/air@latest

Após a instalação, o binário do air estará em ~/go/bin (Linux/macOS) ou %USERPROFILE%\go\bin (Windows).
Adicione esse caminho ao seu PATH para usar o comando globalmente.

3. Rodar o projeto
Dentro da raiz do projeto:

air

Isso iniciará o servidor com auto reload. Toda alteração nos arquivos .go recompila e reinicia o app automaticamente.

⚙️ Exemplo de .air.toml

# Arquivo: .air.toml

tmp_dir = "tmp"
bin = "main.exe"
cmd = "main.go"
🛠 Tecnologias
Go (Golang)

Gin Gonic

GORM

SQLite / MySQL / PostgreSQL (configurável)
