#!/bin/bash

# Verifica se um diretório foi fornecido
if [ $# -eq 0 ]; then
    echo "Por favor, forneça um diretório."
    exit 1
fi

diretorio="$1"
diretorio_destino="/home/$USER/nzb/"

# Verifica se o diretório fornecido existe
if [ ! -d "$diretorio" ]; then
    echo "O diretório fornecido não existe."
    exit 1
fi

# Loop através de todos os arquivos de vídeo no diretório especificado
for arquivo in "$diretorio"/*.{mp4,mkv,avi,mpg,mpeg}; do
    if [ -f "$arquivo" ]; then
        nome_base=$(basename "$arquivo" | sed 's/\.[^.]*$//')
        mediainfo "$arquivo" > "${diretorio}/${nome_base}.nfo"
        mv "${diretorio}/${nome_base}.nfo" "$diretorio_destino"
    fi
done

# Executa ngPost com o diretório fornecido e gera par2
ngPost --auto "$diretorio" --gen_par2
