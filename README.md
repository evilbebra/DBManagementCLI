# DBManagementCLI
CLI tool for management db server

App build:    
<code>docker build -t udvapp . </code>      

В качестве СУБД будет postgresql контейнер с именем pg:    
![postgres](https://github.com/evilbebra/DBManagementCLI/pg.jpg)    
     
Чтобы приложение в докере могло взаимодействовать с контейнером postgresql поместим их    
в докер-сеть:    
<code>docker network create my-network</code>      
<code>docker network connect my-network pg</code>      
<code>docker run --name udv --network=my-network -d udvapp</code>      
    
    
Запустив контейнер, проверим, что приложение работает:  
<code>docker exec -it udv /bin/bash </code> 
<code>./main</code>    
![output1](https://github.com/evilbebra/DBManagementCLI/output1.jpg)      
