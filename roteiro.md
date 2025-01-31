gere um script que receba um arquivo ou um diretório, se for fornecido um arquivo, gere o .nfo deste arquivo utilizando o binário mediainfo e nomeie o arquivo .nfo com o mesmo nome do arquivo de vídeo, depois encaminhe o arquivo .nfo gerado para a pasta ~/nzb. se for fornecido um diretório, gere os .nfo de todos os arquivos dentro do diretório e os trate da mesma maneira.

faça upload do arquivo recebido ou de todos os arquivos contidos no diretório recebido separadamente utilizando o binário ngPost, se não for fornecido nenhum argumento, adicione o argumento --gen_par2 no comando do ngPost, se for fornecido o argumento -p adicione o argumento --pack no comando do ngPost, se for fornecido o argumento -x ofusque a postagem do ngPost

após o fim do processamento do ngPost encerre o programa