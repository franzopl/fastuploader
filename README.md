# fastuploader
Um script para facilitar o uso do ngPost


- Este script foi criado para facilitar o envio de arquivos para usenet.

- É necessário instalar o ngPost antes de executar o script, recomendo instalar o appimage.
  
    wget https://github.com/mbruel/ngPost/releases/download/v4.16/ngPost_v4.16_libssl3-x86_64.AppImage # verifique a melhor versão
    sudo chmod +x ngPost_v4.16_libssl3-x86_64.AppImage # torna o arquivo executável
    sudo mv ngPost_v4.16_libssl3-x86_64.AppImage /usr/local/bin/ngPost # envia ngPost para o PATH do sistema

- Também é necessário instalar o mediainfo.
    sudo apt install mediainfo #ubuntu
    sudo dnf install mediainfo #fedora
    *verifique o gerenciador de pacotes da sua distribuição linux

- o arquivo ngPost.conf deve ser renomeado para .ngPost e colocado na pasta pessoal do usuário.

      mv ./fastuploader/config/ngPost.conf ~/$USER/.ngPost

- Dentro do arquivo .ngPost devem ser editados as seguintes linhas:

      #PROXY_SOCKS5 = proxyuser:proxypassword123@100.100.1.1:1080
      esta configuração é opcional, se você quiser adicionar um proxy SOCKS5 apenas retire o # do início da linha e preencha com os dados do seu proxy.

      GROUPS   = alt.binaries.test
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



Realizada a configuração inicial você pode enviar o script fastuploader para o PATH








