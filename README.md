# Progetto IdS in Go

## Descrizione
Il progetto di IdS è un progetto (di gruppo) che mi è stato assegnato nel corso 
di ingegneria del software. 
In poche parole si trattava di scrivere in programma che permettesse 
di scaricare degli articoli da due fonti: The Guardian e New York Times.
In particolare gli articoli del The Guardian dovevano essere scaricati tramite API;
gli articoli del New York Times venivano fornite tramite un file in formato csv.

Il progetto originale è stato scritto in Java, io ho deciso di prendere tale progetto
e riscriverlo completamente in Go. Le funzionalità offerte dalle due versioni sono
esattamente le stesse, in questa versione ho solamente cercato di migliorare le 
prestazioni del programma sfruttando le strutture dati offerte da Go per la gestione
della sincronizzazione.

Riscrivendo il codice in Go mi sono accorto della semplicità rispetto a Java. 
La cosa che mi ha stupito particolarmente è stata la facilità nell'utilizzare le 
librerie: per utilizzare le librerie di serializzazione e deserializzazione in xml, 
deserializzare da file Json, deserializzare da file csv e per gestire gli argomenti
da linea di comando in Java è stato molto più difficile. 

Successivamente riporto alcune specifiche del progetto e infine riporto una tabella
con i risultati ottenuti.

## I test

I test vengono effettuati variando il numero di articoli che vengono scaricati 
dalle sorgenti. Per ora variano solamente gli articoli richiesti alle API del 
The Guardian, questo perché per ora non ho modo di ottenere ulteriori file csv da 
New York Times. 

Nello specifico sono andato ad raddoppiare ad ogni test il numero di richieste
alle API: entrambi i progetti gestiscono il download sfruttando la concorrenza, 
entrambi i progetti creano thread logici e non fisici. In Java ho utilizzato la
classe `Thread` mentre in Go utilizzo le `Goroutines`.

Al di fuori dei metodi di download degli articoli del The Guardian, il resto del 
codice dovrebbe essere pressoché identico. Testando i due codici senza il download 
del The Guardian, quindi con solo il "dowload" degli articoli del New York Times 
e la successiva serializzazione, deserializzazione e conteggio parole, si nota che
la versione in Go è quasi il doppio più veloce. (DA RICONTROLLARE CON TEST: inoltre
modifica il codice gestione csv di Java).
In ogni caso si tratta di un tempo costante.

Dunque in conclusione dovremmo ottenere un test (non particolarmente preciso) 
che dipende dalla gestione dei thread logici effettuata dai due linguaggi.

Si tratta di stime grossolane che potrebbero non essere del tutto veritiere visto 
che il codice non è esattamente identico in entrambi i progetti. In ogni caso
si tratta di un progett realizzato per pura noia che non ha grandi pretese.

# Risultati

Le voci:
- N test: numero di test totali effettuati.
- Articoli: totale di articoli scaricati.
- Tempo totale: somma dei tempi di esecuzione dei test.
- Tempo medio: tempo medio test (Tempo totale/N test)

## Fase di download

Gli articoli scaricati dal The Guardian sono sempre 1000; variano quelli del 
New York Times, avendo a disposizione un unico file csv ho utilizzato solo quello.

Dal numero totale di articoli si può risalire al numeoro di thread (`Thread` e 
`Goroutines`) lanciati. I thread lanciati per scaricare gli articoli del The Guardian
sono sempre 5, non li considero nel seguenti calcoli:

Ogni thread gestisce 1000 articoli. Il numero di thread è dato dalla formula 
(articoli-1000)/1000.

### Go

|N test | Articoli  |Tempo totale | Tempo medio |  
|-------|-----------|-------------|-------------|
| 6     | 2001      |   13.41     |   2.235     |  
| 6     | 6005      |   18.77     |   3.128     |  
| 6     | 11010     |   17.22     |   2.87      |  
| 6     | 21020     |   21.23     |   3.538     |  
| 6     | 41040     |   37.96     |   6.326     |  
| 6     | 81080     |   39.29     |   6.548     |  
| 6     | 161160    |   57.22     |   9.536     |  
| 6     | 361360    |   79.44     |   13.24     |  
| 6     | 501500    |   150.86    |   25.143    |  

Il numero massimo massimo di articoli gestibile con il metodo che ho scritto è 
di circa 524 mila. Oltre tale soglia viene lanciato un'errore che dice "too many 
open files". Adottando un'altra soluzione penso si possa aggirare tale vincolo.

### Java

|N test | Articoli  | Tempo totale | Tempo medio |  
|-------|-----------|--------------|-------------|
| 6     | 2000      |    22.64     |   3.773     |  
| 6     | 6000      |    26.52     |   4.42      |  
| 6     | 11000     |    30.37     |   5.061     |  
| 6     | 21000     |    41.76     |   6.96      |  
| 6     | 41000     |    80.6      |   13.433    |  
| 6     | 81000     |    82.05     |   13.675    |  
| 6     | 161000    |    147.23    |   24.538    |  
| 6     | 361000    |    293.83    |   48.971    |  

Il numero massimo di articoli gestibile con il metodo scritto è circa 190 mila. 
Oltre a tale soglia viene generato l'errore `OutOfMemoryError` dovuto al Java Heap
Space. Anche in questo caso penso che il vincolo sia una conseguenza di come ho 
scritto il metodo.

### Note sul dowload

Java utilizzava molta più cpu rispetto a Go.

## Fase di estrazione

Entrambi gli algoritimi non sfruttano la concorrenza, infatti le prestazioni non sono
ottime, penso che potrei fare di megio. Ho testato le prestazioni di questo algoritmo
perché si tratta della strategia adottata nel progetto originale.

### Go

|N test | Articoli  | Tempo totale | Tempo medio |  
|-------|-----------|--------------|-------------|
| 6     | 2001      |    3.17      |   0.528     |  
| 6     | 6005      |    6.48      |   1.08      |  
| 6     | 11010     |    10.46     |   1.743     |  
| 6     | 21020     |    18.4      |   3.066     |  
| 6     | 41040     |    34.1      |   5.6833    |  
| 6     | 81080     |    65.49     |   10.915    |  
| 6     | 161160    |    129.01    |   21.50     |  
| 6     | 361360    |    260.37    |   43.39     |  
| 6     | 501500    |    652.5     |   108.75    |  

### Java

|N test | Articoli  | Tempo totale | Tempo medio |  
|-------|-----------|--------------|-------------|
| 6     | 2000      |   5.25       |   0.875     |  
| 6     | 6000      |   9.97       |   1.661     |  
| 6     | 11000     |   15.89      |   2.648     |  
| 6     | 21000     |   25.07      |   4.178     |  
| 6     | 41000     |   41.26      |   6.876     |  
| 6     | 81000     |   73.81      |   12.301    |  
| 6     | 161000    |   156.19     |   26.031    |  
| 6     | 361000    |   201.55     |   33.591    |  

