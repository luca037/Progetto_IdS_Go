# Progetto IdS in Go

## Descrizione
Il progetto di IdS è un progetto (di gruppo) che mi è stato assegnato nel corso 
di ingegneria del software. 
In poche parole si trattava di scrivere in programma che permettesse 
di scaricare degli articoli da due fonti: The Guardian e New York Times.
In particolare gli articoli del The Guardian devevano essere scaricati tramite API;
gli articoli del New York Times venivano fornite tramite un file in formato csv.

Il progetto originale è stato scritto in Java, io ho deciso di prendere tale progetto
e riscriverlo completamente in Go. Le funzionalità offerte dalle due versioni sono
esattamente le stesse, in questa versione ho solamente cercato di migliorare le 
prestazioni del programma sfruttando le strutture dati offerte da Go per la gestione
della sincronizzazione.

Riscrivendo il codice in Go mi sono accorto della semplicità rispetto a Java. 
La cosa che mi ha stupito particolarmente è stata la facilità nell'utilizzare le 
librerie: per utilizzare le librerie di serializzazione/deserializzazione in xml, 
deserializzare da file Json, deserializzare da file csv e per gestire gli argomenti
da linea di comando in Java è stato molto più difficile, ho speso decisamente più tempo
rispetto a quanto ne ho speso nella versione in Go. 

Successivamante riporto alcune specifiche del progetto e infine riporto una tabella
con i risultati ottenuti.
