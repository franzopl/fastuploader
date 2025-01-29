#!/bin/bash

# Muda para o diretório correto
cd /home/$USER/nzb/ || exit

# Certifique-se de que a pasta 'enviados' existe
mkdir -p enviados

# Variável para armazenar a chave da API (substitua com sua chave real)
myapikey="koxEbgbGUPnBOfBcPV6taHl3VuxhwSPR"

# Loop para processar cada arquivo .nzb
for nzb_file in *.nzb; do
    if [ -f "$nzb_file" ]; then
        # Encontra o arquivo .nfo correspondente
        nfo_file="${nzb_file%.nzb}.nfo"
        
        if [ -f "$nfo_file" ]; then
            # Executa o comando curl para enviar os arquivos
            response=$(curl -k -s -L -m 60 -F "nzb=@$nzb_file" -F "nfo=@$nfo_file" "https://api.nzbgeek.info/submit?apikey=$myapikey")
            
            # Verifica se a resposta indica sucesso
            if [[ $response == *"{\"response\":{\"@attributes\":{\"API\":\"OK\",\"NFO\":\"OK\",\"REGISTER\":\"OK\"}}}"* ]]; then
                echo "Envio bem-sucedido para $nzb_file e $nfo_file"
                mv "$nzb_file" "$nfo_file" enviados/
            else
                echo "Falha no envio para $nzb_file e $nfo_file. Resposta: $response"
            fi
        else
            echo "Arquivo .nfo correspondente ($nfo_file) não encontrado para $nzb_file"
        fi
    fi
done

echo "Processamento concluído."
