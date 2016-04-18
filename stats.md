# Práctica 1.1

- Nº de ficheros procesados: `ls -1 *.tok`
- Nº de tokens totales: `cat *.tok | wc -l`

# Práctica 1.2

- Nº de tokens únicos antes del stoper: `cat *.tok | sort | uniq | wc -l`
- Nº de tokens únicos después del stopper: `cat *.stop | sort | uniq | wc -l`
- 5 palabras más repetidas antes del stopper: `cat *.tok | sort | uniq -c | sort -nr | head -n 5`
- 5 palabras más repetidas después del stopper: `cat *.stop | sort | uniq -c | sort -nr | head -n 5`
- Máximo número de tokens por fichero: `wc -l *.tok | sort -nr | tail -n+2 | head -n 1`
- Mínimo número de tokens por fichero `wc -l *.tok | sort -n | head -n 1`

# Práctica 1.3
- Nº de tokens únicos después del stemmer: `cat *.stem | sort | uniq | wc -l`

- 5 palabras más repetidas antes del stemmer: `cat *.stem | sort | uniq -c | sort -nr | head -n 5`
- 5 palabras más repetidas después del stemmer: `cat *.stem | sort | uniq -c | sort -nr | head -n 5`
- Máximo número de tokens por fichero: `wc -l *.tok | sort -nr | tail -n+2 | head -n 1`
- Mínimo número de tokens por fichero `wc -l *.tok | sort -n | head -n 1`
