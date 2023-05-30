#!/bin/bash

g++ -o compilado *.cpp

if [ $? -eq 0 ];then 
        
    valor=0;
   for ((i=1000; i<=100000; i+=1000))
    do    
    echo "compilado com sucesso o valor - $i"
    ./compilado $i >> resul
    done           
fi
