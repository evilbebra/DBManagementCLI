# DBManagementCLI
CLI tool for management db server

## App build:    
<code>docker build -t udvapp . </code>      

В качестве СУБД будет postgresql контейнер с именем pg:    
POSTGRES_PASSWORD=docker    
POSTGRES_USER=docker    
![postgres](https://raw.githubusercontent.com/evilbebra/DBManagementCLI/master/readmd-images/pg.jpg)    
     
Чтобы приложение в докере могло взаимодействовать с контейнером postgresql поместим их    
в докер-сеть:    
<code>docker network create my-network</code>      
<code>docker network connect my-network pg</code>      
<code>docker run --name udv --network=my-network -d udvapp</code>      
    
    
Запустив контейнер, проверим, что приложение работает:  
<code>docker exec -it udv /bin/bash </code>    
<code>./main</code>    
![output1](https://raw.githubusercontent.com/evilbebra/DBManagementCLI/master/readmd-images/output1.jpg)    
    
Также создание бэкапа и восстановление БД:    
![create-table](https://raw.githubusercontent.com/evilbebra/DBManagementCLI/master/readmd-images/create-table.jpg)
![backup1](https://raw.githubusercontent.com/evilbebra/DBManagementCLI/master/readmd-images/backup1.jpg)    
![nano](https://raw.githubusercontent.com/evilbebra/DBManagementCLI/master/readmd-images/nano.jpg)
![restore](https://raw.githubusercontent.com/evilbebra/DBManagementCLI/master/readmd-images/restore.jpg)
