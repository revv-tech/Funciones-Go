"""
Generar número pseudoaleatorios
Adaptado de Dromey, 'How to solve it by computer'
"""

def aleatorio (x, n):
    ## x es la semilla
    ## n es la cantidad de números pseudoaleatorios que deseamos generar

    ## inicializaciones
    m = 2048 ## período
    a = 53  ## multiplicador
    b = 541  ## incremento

    for i in range (n):
        ## calcular el siguiente número aleatorio
        x = (a * x + b) % m

        ## devolver el siguiente número aleatorio
        OldRange = (2048 - 0)  
        NewRange = (199 - 0)  
        x = int((((x - 0) * NewRange) / OldRange) + 0)
        yield x

    ## terminar (el for consumidor no llegará aquí)
    return "fin"

## Semilla
semilla = int(input ("Ingrese una semilla, por favor: "))
cuántos = int(input ("¿Cuántos números pseudo-aleatorios desea generar?: "))

for r in aleatorio (semilla, cuántos):
    print (r)

lista = [ r for r in aleatorio (semilla, cuántos) ]

    
        

    
    
