 #!/bin/sh
 awk 'NR == FNR { if(FNR >= 17 && FNR <= 23) { patch = patch $0 ORS }; next } FNR == 2 { $0 = patch $0 } 1' native/privatebank.go js/empty.go.js > js/privatebank_abibin.go