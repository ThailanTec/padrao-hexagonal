## Organização de pastas


### Internal

Logica interna do projeto, não pode ser importada, devendo ficar apenas ali dentro. 

### pkg

Como se fosse uma lib, para "servir" no projeto. 

### cmd

Arquivos de inicialização da aplicação e outros. 


### Raiz

Pode deixar o main, go.mod e assim utilizar de uma forma mais "clean", deixando o git.ignore, docker e outros. 

Ordem:

Domain/Ports > service > repository > handle... 