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

        - verifique o gerenciador de pacotes da sua distribuição linux

 - Também é necessário instalar o rar
    
       sudo apt install rar #ubuntu

 - verifique o gerenciador de pacotes da sua distribuição linux ou baixe o binário em https://www.win-rar.com/fileadmin/winrar-versions/rarlinux-x64-701.tar.gz extraia o tar.gz e envie o binário rar para o PATH

       wget https://www.win-rar.com/fileadmin/winrar-versions/rarlinux-x64-701.tar.gz
   
       tar -xvzf rarlinux-x64-701.tar.gz
   
       sudo chmod +x ./rar/rar
   
       sudo ln ./rar/rar /user/bin/rar  
    
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

      sudo chmod +x ./scripts/fastupload
      sudo ln ./scripts/fastupload /usr/bin/fastupload


após instalado no PATH basta realizar os seguintes comandos para realizar o upload:

      fastupload <arquivo ou diretório>
      exemplo: 
        fastupload teste.mkv #irá gerar arquivo par2 e enviar o arquivo teste.mkv, esta é a configuração padrão
        fastupload ./teste #ao indicar um diretório irá gerar o par2 de cada arquivo e enviar separadamente, a forma mais fácil de enviar todos os arquivos de uma pasta é navegar até a pasta e digitar "fastupload . " (o ponto indica que o diretório atual será selecionado)


há 2 argumentos que podem ser utilizados para postagens com password ou ofuscadas

      adicione -p ao comando para ativar a função --pack do ngPost que irá comprimir o arquivo, embaralhar o nome, adicionar uma senha e gerar o par2 antes de enviar.  
      adiciona -x para ofuscar o cabeçalho da postagem.  
      ex: fastupload -p video.mkv
          fastupload -p -x video.mkv
      
Lembre-se que ao utiliar essas opções você dificilmente irá recuperar sem o arquivo .nzb, utilize com cuidado essas opções pois postagens ofuscadas em excesso podem ser consideradas SPAM pelo servidor.


Para opções de upload mais elaboradas você pode editar o arquivo .ngPost e utilizar os comandos diretamente no ngPost.




