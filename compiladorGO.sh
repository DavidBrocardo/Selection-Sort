#!/bin/bash

valor=0;
for ((i=1000; i<=100000; i+=1000))
 do    
  echo "compilado com sucesso o valor - $i"
  go run *.go $i >> resul
  
done           

