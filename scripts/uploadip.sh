#!/bin/bash

# Verifica se um arquivo foi fornecido
if [ $# -eq 0 ]; then
    echo "Por favor, forneça um arquivo."
    exit 1
fi

arquivo="$1"
nome_base=$(basename "$arquivo" | sed 's/\.[^.]*$//')
diretorio_destino="/home/$USER/nzb/"

# Gera o arquivo .nfo usando mediainfo
mediainfo "$arquivo" > "${nome_base}.nfo"

# Move o arquivo .nfo para o diretório especificado
mv "${nome_base}.nfo" "$diretorio_destino"

# Executa ngPost com o arquivo fornecido e gera par2
ngPost -i "$arquivo" --pack
