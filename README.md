# ✈️ itinerary-prettifier

A command-line tool that prettifies flight itineraries for customers.  
It reads a text-based itinerary, processes airport codes, dates, and formatting, and writes a customer-friendly version to a new file.

---

## 🧭 Usage

```bash
go run . <input_file> <output_file> <airport_lookup.csv>
```

Example:
```bash
go run . ./input.txt ./output.txt ./airport-lookup.csv
```

Show help:
```bash
go run . -h
```

Output:
```
itinerary usage:
go run . ./input.txt ./output.txt ./airport-lookup.csv
```

---

## 📄 Functionality

- Converts airport codes:
  - `#LAX` → Los Angeles International Airport (IATA)
  - `##EGLL` → London Heathrow Airport (ICAO)
  - `*#LHR` → London (city name)
- Formats dates and times:
  - `D(2007-04-05T12:30−02:00)` → `05 Apr 2007`
  - `T12(2007-04-05T12:30−02:00)` → `12:30PM (-02:00)`
  - `T24(2007-04-05T12:30−02:00)` → `12:30 (-02:00)`
- Trims extra blank lines and normalizes line breaks (`\v`, `\f`, `\r` → `\n`).

---

## 🧾 Airport Lookup

The lookup CSV must have the following columns **in this order**:

```
name, iso_country, municipality, icao_code, iata_code, coordinates
```

If any column is missing or empty, the program prints:

```
Airport lookup malformed
```

---

## ⚠️ Error Handling

| Condition | Message |
|------------|----------|
| Missing arguments | `itinerary usage: go run . ./input.txt ./output.txt ./airport-lookup.csv` |
| Input not found | `Input not found` |
| Airport lookup missing | `Airport lookup not found` |
| Malformed lookup file | `Airport lookup malformed` |

In all error cases, the output file is **not created or modified**.

---

## 🌟 Bonus Features

- City names supported (e.g. `*#LHR → London`)  
- Output also printed to standard output (stdout)

---

## 🧪 Example Test

### Example input (`input.txt`)
```
Flight itinerary for customer

Departing from #LAX to ##EGLL
Date: D(2025-04-05T12:30-02:00)
Departure time: T24(2025-04-05T12:30-02:00)
Arrival time: T12(2025-04-05T23:30Z)

Connecting flight from *#CDG to #XXX

\v\f

Thank you for flying with us!

here I test
 \r


here I test  \v


here I test  \l
```

### Example airport lookup (`airport-lookup.csv`)
```
name,iso_country,municipality,icao_code,iata_code,coordinates
Los Angeles International Airport,US,Los Angeles,KLAX,LAX,33.9425,-118.4081
London Heathrow Airport,GB,London,EGLL,LHR,51.4775,-0.4614
Charles de Gaulle International Airport,FR,Paris,LFPG,CDG,49.0097,2.5479
```

### Expected output (`output.txt`)
```
Flight itinerary for customer

Departing from Los Angeles International Airport to London Heathrow Airport
Date: 05 Apr 2025
Departure time: 12:30 (-02:00)
Arrival time: 11:30PM (+00:00)

Connecting flight from Paris to #XXX

Thank you for flying with us!

here I test 
 \r

here I test 
 \v

here I test 
 \l
```

---

## 📚 Learning Outcomes

- Command-line argument handling  
- File reading/writing in Go  
- String manipulation and parsing  
- Working with CSV data and error validation  

---

**Author:** Bruno Zambrano  
[https://gitea.kood.tech/brunozambrano/itinerary](https://gitea.kood.tech/brunozambrano/itinerary)
