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
The Guardian, questo perché per ora non ho modo di ottere ulteriori file csv da 
New York Times. 

Nello specifico sono andato ad raddoppiare ad ogni test il numero di richieste
alle API: entrambi i progetti gestiscono il download sfruttando la concorrenza, 
entrambi i progetti creano thread logici e non fisici. In Java ho utilizzato la
classe `Thread` mentre in Go utilizzo le `Goroutines`.

Al di fuori dei metodi di download degli articoli del The Guardian, il resto del 
codice dovrebbe essere pressoché idendico. Testando i due codici senza il download 
del The Guardian, quindi con solo il "dowload" degli articoli del New York Times 
e la successiva serializzazione, deserializzazione e conteggio parole, si nota che
la versione in Go è quasi il doppio più veloce. (DA RICONTROLLARE CON TEST: inoltre
modifica il codice gestione csv di Java).
In ogni caso si tratta di un tempo costante.

Dunque in conclusione dovremmo ottenere un test (non particolarmente preciso) 
che dipende dalla gestione dei thread logici effettuata dai due linguaggi.

Si tratta di stime grossolane che potrebbero non essere del tutto veritiere visto 
che il codice non è esattamente identico in entrambi i progetti. In ogni caso
si tratta di un progett realizzato per pura noia.
