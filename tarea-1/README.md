# Laboratorio 2 - Sistemas Distribuidos

## Grupo 2
- Leandro Espinoza 201973563-2
- Miguel Soto 201973623-K

## Instrucciones

Instrucciones de uso:
Ejecutar servidor central con "make docker" y ejecutar capitanes con "make capitanes"


## Si se agregara una central que ejecute las mismas funciones que la actualmente implemen-
tada. ¿Esto beneficiar´ıa o complicar´ıa al proceso?

No, de hecho generarian inconsistencias en los datos, ya que los capitanes estarian
mandando mensajes entre ambas centrales, las cuales manejarian datos diferentes de
los botines,