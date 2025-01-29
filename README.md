# fastuploader
scripts for usenet posting


- Este script foi criado para facilitar o envio de arquivos para usenet.

- É necessário instalar o ngPost e o mediainfo antes de executar o script.

- o arquivo ngPost.conf deve ser renomeado para .ngPost e colocado na pasta pessoal do usuário.

```bash
  mv ./fastuploader/config/ngPost.conf ~/$USER/.ngPost
```
- Dentro do arquivo .ngPost devem ser etidados as seguintes linhas:

``` bash
#PROXY_SOCKS5 = proxyuser:proxypassword123@100.100.1.1:1080
```
    esta configuração é opcional, se você quiser adicionar um proxy SOCKS5 apenas retire o # do início da linha e preencha com os dados do seu proxy.


#GROUPS   = alt.binaries.test

    Altere aqui o newsgroup para o qual deseja fazer seus uploads

RAR_PATH = /usr/bin/rar
    
    preencha aqui o caminho para o seu binário rar


[server]
host = usnews.blocknews.net  
port = 563  
ssl  = true  
user =  
pass =  
connection = 40  
enabled = true  
nzbCheck = false  

    preencha aqui os dados do seu servidor

Os outros dados podem ser adaptados a sua necessidade seguindo a configuração do ngPost, porém não é necessário.








