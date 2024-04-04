# Laboratorio 2 - Sistemas Distribuidos

## Grupo 2
- Leandro Espinoza 201973563-2
- Miguel Soto 201973623-K

## Instrucciones

Instrucciones de uso:
Ejecutar servidor central con "make docker" y ejecutar capitanes con "make capitanes"


## Si se agregara una central que ejecute las mismas funciones que la actualmente implemen-
tada. ¿Esto beneficiar´ıa o complicar´ıa al proceso?

Si, en base a lo que hemos visto en clase, una central extra actuaria como un balanceador
de carga para los clientes (capitanes), por lo que si uno de estos se ve con mucha carga
los clientes podrian acceder a la central nueva para acelerar el proceso.